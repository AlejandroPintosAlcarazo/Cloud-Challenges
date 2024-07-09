//go:build setup
// +build setup

package main

import (
	"context"
	"log"

	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/fetcher"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/load"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/query"
)

func main() {
	ctx := context.Background()

	client, err := query.ConnectBigQuery(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to BigQuery: %v", err)
	}

	if err := query.SetupBigQuery(ctx, client); err != nil {
		log.Fatalf("Failed to setup BigQuery: %v", err)
	}

	simpleStations, err := fetcher.FetchStationsData()
	if err != nil {
		log.Fatalf("Failed to fetch all station IDs: %v", err)
	}

	for _, simpleStation := range simpleStations {
		if err := load.LoadStation(ctx, client, simpleStation); err != nil {
			log.Printf("Failed to process and save station %s: %v", simpleStation.ID, err)
		} else {
			log.Printf("Successfully processed and saved station %s", simpleStation.ID)
		}
	}

	log.Println("Station data uploaded to BigQuery completed!")
}
