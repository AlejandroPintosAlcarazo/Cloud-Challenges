package load

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/configs"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/models"
)

func saveStation(ctx context.Context, client *bigquery.Client, station models.Station) error {
	inserter := client.Dataset(configs.DatasetID).Table(configs.TableStations).Inserter()

	// Imprimir datos de la estación
	//fmt.Printf("Inserting station data: %+v\n", station)

	if err := inserter.Put(ctx, station); err != nil {
		return fmt.Errorf("failed to insert station data: %v", err)
	}
	return nil
}

func saveWeatherData(ctx context.Context, client *bigquery.Client, data *models.DataEntry) error {
	inserter := client.Dataset(configs.DatasetID).Table(configs.TableWeather).Inserter()

	// Imprimir datos de la estación
	//fmt.Printf("Inserting weather data: %+v\n", data)

	if err := inserter.Put(ctx, data); err != nil {
		return fmt.Errorf("failed to insert weather data: %v", err)
	}
	return nil
}

func saveError(ctx context.Context, client *bigquery.Client, problemDate, tableName, stationID, message string) error {
	inserter := client.Dataset(configs.DatasetID).Table(configs.TableErrors).Inserter()

	errorRecord := models.ServerErrorRecord{
		ProblemDate: problemDate,
		Table:       tableName,
		StationID:   stationID,
		Message:     message,
		Timestamp:   time.Now(),
		Fixed:       false,
	}

	if err := inserter.Put(ctx, errorRecord); err != nil {
		return fmt.Errorf("failed to insert error record: %v", err)
	}
	return nil
}
