package utils

import (
	"fmt"
	"sort"
	"time"

	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/configs"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/models"
)

func CreateMappingData(chunk [2]string, data []models.DataEntry) ([]string, []string) {
	start := chunk[0]
	end := chunk[1]

	// Crear un mapa de todas las fechas en el chunk
	dateMap := make(map[string]struct{})
	current, _ := time.Parse("2006-01-02", start)
	endDate, _ := time.Parse("2006-01-02", end)
	for !current.After(endDate) {
		dateMap[current.Format("2006-01-02")] = struct{}{}
		current = current.AddDate(0, 0, 1)
	}

	// Crear un array para las fechas presentes en data y otro para las faltantes
	var presentDates []string
	var missingDates []string

	// Marcar las fechas presentes en data
	for _, entry := range data {
		if _, exists := dateMap[entry.Date]; exists {
			presentDates = append(presentDates, entry.Date)
			delete(dateMap, entry.Date)
		}
	}

	// Las fechas restantes en dateMap son las faltantes
	for date := range dateMap {
		missingDates = append(missingDates, date)
	}

	// Ordenar las fechas en orden creciente
	sort.Strings(presentDates)
	sort.Strings(missingDates)
	return presentDates, missingDates
}

func CreateChunks(startDate, endDate string) ([][2]string, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %v", err)
	}
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %v", err)
	}

	var chunks [][2]string

	for current := start; current.Before(end); {
		chunkEnd := current.AddDate(0, 0, configs.ChunkSize-1)
		if chunkEnd.After(end) {
			chunkEnd = end
		}
		chunks = append(chunks, [2]string{
			current.Format("2006-01-02"),
			chunkEnd.Format("2006-01-02"),
		})
		current = chunkEnd.AddDate(0, 0, 1)
	}

	return chunks, nil
}
