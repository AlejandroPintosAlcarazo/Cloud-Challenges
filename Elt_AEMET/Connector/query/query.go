package query

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/configs"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/models"
	"google.golang.org/api/iterator"
)

func QueryStationInfo(ctx context.Context, client *bigquery.Client) ([]models.StationUpdate, error) {
	//Crear la Query
	query := fmt.Sprintf(`
		SELECT id, final,
		FROM %s.%s
	`, configs.DatasetID, configs.TableStations)

	//Lanzar la query
	it, err := client.Query(query).Read(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	//Guardar la query en variables
	var stations []models.StationUpdate
	for {
		var station models.StationUpdate
		err := it.Next(&station)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to iterate over query results: %v", err)
		}
		stations = append(stations, station)
	}

	return stations, nil
}

func CheckForStreamingBuffer(ctx context.Context, client *bigquery.Client) error {
	// Verificar la tabla Station
	table := client.Dataset(configs.DatasetID).Table(configs.TableStations)
	meta, err := table.Metadata(ctx)
	if err != nil {
		return fmt.Errorf("failed to get table metadata: %v", err)
	}

	// Check if there are rows in the streaming buffer
	if meta.StreamingBuffer != nil {
		fmt.Println("Streaming buffer detected, waiting for it to clear...")
		time.Sleep(1 * time.Minute)
		return CheckForStreamingBuffer(ctx, client)
	}

	// Verificar la tabla Weather
	tableWeather := client.Dataset(configs.DatasetID).Table(configs.TableWeather)
	metaWeather, err := tableWeather.Metadata(ctx)
	if err != nil {
		return fmt.Errorf("failed to get table metadata for Weather: %v", err)
	}

	// Check if there are rows in the streaming buffer for Weather
	if metaWeather.StreamingBuffer != nil {
		fmt.Println("Streaming buffer detected for Weather, waiting for it to clear...")
		time.Sleep(1 * time.Minute)
		return CheckForStreamingBuffer(ctx, client)
	}
	return nil
}
