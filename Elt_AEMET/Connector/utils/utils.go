package utils

import (
	"fmt"
	"time"

	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/configs"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/fetcher"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/models"
)

func StationNeedsUpdate(station models.StationUpdate) bool {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	return station.Final < yesterday
}

func WrapInQuotes(strs []string) []string {
	wrapped := make([]string, len(strs))
	for i, s := range strs {
		wrapped[i] = fmt.Sprintf("'%s'", s)
	}
	return wrapped
}

func FindOldestDate(stationID string) (string, error) {
	// Calcular la fecha de finalización con el retraso configurado
	end := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	// Crear los chunks de fechas
	chunks, err := CreateChunks(configs.StartDate, end)
	if err != nil {
		return "", fmt.Errorf("failed to create date chunks: %v", err)
	}

	var oldestDate string

	// Iterar sobre cada chunk de fechas
	for _, chunk := range chunks {
		data, err := fetcher.FetchAEMETData(chunk[0], chunk[1], stationID)
		if err != nil {
			fmt.Printf("Error fetching data for station %s from %s to %s: %v\n", stationID, chunk[0], chunk[1], err)
			continue
		}

		// Verificar si se obtuvieron datos
		if data != nil && len(*data) > 0 {
			// Obtener la primera entrada válida
			oldestDate = (*data)[0].Date
			break
		}
	}

	if oldestDate == "" {
		return "", fmt.Errorf("no data found for station %s", stationID)
	}

	return oldestDate, nil
}
