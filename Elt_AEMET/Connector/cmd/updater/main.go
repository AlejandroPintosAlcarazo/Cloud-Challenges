//go:build connector
// +build connector

package main

import (
	"context"
	"log"

	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/load"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/models"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/query"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/utils"
)

func main() {
	go utils.HandleCloudRun()
	ctx := context.Background()

	// Conectar a BigQuery
	client, err := query.ConnectBigQuery(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to BigQuery: %v", err)
	}

	// Obtener estaciones
	stations, err := query.QueryStationInfo(ctx, client)
	if err != nil {
		log.Fatalf("Failed to query station info: %v", err)
	}

	// Esperar a que el streaming buffer se vacíe antes de actualizar en masa
	if err := query.CheckForStreamingBuffer(ctx, client); err != nil {
		log.Fatalf("Failed to wait for streaming buffer to clear: %v", err)
	}

	// Crear un slice para almacenar los StationState
	var stationStates []models.StationState

	// Iterar sobre todas las estaciones
	for _, station := range stations {
		if utils.StationNeedsUpdate(station) {
			stationState, err := load.LoadWeather(ctx, client, station)
			if err != nil {
				log.Printf("Failed to fetched station data %s: %v", station.ID, err)
			} else {
				log.Printf("Successfully fetched station data %s", station.ID)
				stationStates = append(stationStates, stationState)
			}
		} else {
			log.Printf("No update needed for station %s", station.ID)
		}
	}

	//Actualizar el estado de las estaciones
	if err := load.LoadStationState(ctx, client, stationStates); err != nil {
		log.Fatalf("Failed to update stations state: %v", err)
	} else {
		log.Println("Successfully updated all stations states")
	}
}
