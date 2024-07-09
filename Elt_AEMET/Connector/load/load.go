package load

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/configs"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/fetcher"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/models"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/utils"
)

func LoadStation(ctx context.Context, client *bigquery.Client, simpleStation models.SimpleStation) error {
	station := models.Station{
		ID:        simpleStation.ID,
		Nombre:    simpleStation.Nombre,
		Provincia: simpleStation.Provincia,
		Latitud:   simpleStation.Latitud,
		Longitud:  simpleStation.Longitud,
	}

	oldestDate, err := utils.FindOldestDate(simpleStation.ID)
	if err != nil {
		return fmt.Errorf("failed to find oldest date for station %s: %v", simpleStation.ID, err)
	}

	// Actualizar los datos de la estación con la fecha más antigua y las fechas faltantes
	station.Inicio = oldestDate
	station.Final = oldestDate

	// Guardar la estación en BigQuery
	if err := saveStation(ctx, client, station); err != nil {
		saveError(ctx, client, oldestDate, configs.TableStations, station.ID, err.Error())
		return fmt.Errorf("failed to save station data: %v", err)
	}

	return nil
}

func LoadWeather(ctx context.Context, client *bigquery.Client, station models.StationUpdate) (models.StationState, error) {
	// Calcular la fecha final
	end := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	// Crear los chunks de fechas
	chunks, err := utils.CreateChunks(station.Final, end)
	if err != nil {
		return models.StationState{}, fmt.Errorf("failed to create date chunks: %v", err)
	}

	var allPresentDates, allMissingDates []string

	// Iterar sobre cada chunk de fechas
	for _, chunk := range chunks {

		data, err := fetcher.FetchAEMETData(chunk[0], chunk[1], station.ID)
		if err != nil {
			fmt.Printf("Error fetching data for station %s from %s to %s: %v\n", station.ID, chunk[0], chunk[1], err)
			continue
		}

		// Crear el mapeo de fechas existentes y faltantes
		var presentDates, missingDates []string
		if data != nil {
			presentDates, missingDates = utils.CreateMappingData(chunk, *data)
			for _, entry := range *data {
				if err := saveWeatherData(ctx, client, &entry); err != nil {
					log.Printf("Failed to save weather data for %s: %v", entry.Date, err)
					saveError(ctx, client, entry.Date, configs.TableWeather, station.ID, err.Error())
				}
			}
		} else {
			presentDates, missingDates = utils.CreateMappingData(chunk, []models.DataEntry{})
		}
		allMissingDates = append(allMissingDates, missingDates...)

		allPresentDates = append(allPresentDates, presentDates...)
	}

	var lastUpdate string
	if len(allPresentDates) > 0 {
		lastUpdate = allPresentDates[len(allPresentDates)-1]
	} else {
		lastUpdate = station.Final
	}

	// Filtrar allMissingDates para quitar las fechas posteriores a lastUpdate
	var filteredMissingDates []string
	for _, date := range allMissingDates {
		if date <= lastUpdate {
			filteredMissingDates = append(filteredMissingDates, date)
		}
	}
	allMissingDates = filteredMissingDates

	fmt.Printf("\nestacion: %s  y sus fechas faltantes: %v\n", station.ID, filteredMissingDates)
	fmt.Printf("\nestacion: %s  y sus fechas presentes: %v\n", station.ID, allPresentDates)

	stationState := models.StationState{
		ID:        station.ID,
		Final:     lastUpdate,
		Faltantes: filteredMissingDates,
	}

	return stationState, nil
}

func LoadStationState(ctx context.Context, client *bigquery.Client, updates []models.StationState) error {
	updateFinalCases := ""
	updateFaltantesCases := ""
	ids := ""

	for _, update := range updates {
		faltantesStr := fmt.Sprintf("[%s]", strings.Join(utils.WrapInQuotes(update.Faltantes), ", "))
		updateFinalCases += fmt.Sprintf("WHEN id = '%s' THEN '%s' ", update.ID, update.Final)
		updateFaltantesCases += fmt.Sprintf("WHEN id = '%s' THEN ARRAY_CONCAT(faltantes, %s) ", update.ID, faltantesStr)
		ids += fmt.Sprintf("'%s',", update.ID) // Collect IDs for WHERE clause
	}

	// Remove the trailing comma from the ids string
	if len(ids) > 0 {
		ids = ids[:len(ids)-1]
	}

	queryString := fmt.Sprintf(`
        UPDATE %s.%s.%s
        SET 
            final = CASE %s ELSE final END,
            faltantes = CASE %s ELSE faltantes END
        WHERE id IN (%s)
    `, configs.ProjectID, configs.DatasetID, configs.TableStations, updateFinalCases, updateFaltantesCases, ids)

	//Debuggin
	fmt.Println("Constructed Query String:")
	fmt.Println(queryString)

	//Lanzar Query
	q := client.Query(queryString)

	job, err := q.Run(ctx)
	if err != nil {
		saveError(ctx, client, time.Now().Format("2006-01-02"), "stations", "", err.Error())
		return fmt.Errorf("failed to start query job: %v", err)
	}

	status, err := job.Wait(ctx)
	if err != nil {
		saveError(ctx, client, time.Now().Format("2006-01-02"), "stations", "", err.Error())
		return fmt.Errorf("failed to wait for query job completion: %v", err)
	}
	if status.Err() != nil {
		saveError(ctx, client, time.Now().Format("2006-01-02"), "stations", "", status.Err().Error())
		return fmt.Errorf("query job completed with error: %v", status.Err())
	}

	return nil
}
