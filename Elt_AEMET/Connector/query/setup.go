package query

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/configs"
	"google.golang.org/api/option"
)

func ConnectBigQuery(ctx context.Context) (*bigquery.Client, error) {
	client, err := bigquery.NewClient(ctx, configs.ProjectID, option.WithCredentialsFile(configs.CredPath))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	return client, nil
}

func SetupBigQuery(ctx context.Context, client *bigquery.Client) error {
	if err := createDataset(ctx, client); err != nil {
		return err
	}
	if err := createStationTable(ctx, client); err != nil {
		return err
	}

	if err := createWeatherDataTable(ctx, client); err != nil {
		return err
	}

	if err := createErrorTable(ctx, client); err != nil {
		return err
	}

	// Esperar a que el streaming buffer se vacíe antes de actualizar en masa
	if err := CheckForStreamingBuffer(ctx, client); err != nil {
		log.Fatalf("Failed to wait for streaming buffer to clear: %v", err)
	}

	return nil
}
