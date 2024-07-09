package query

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/googleapi"

	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/configs"
)

func createDataset(ctx context.Context, client *bigquery.Client) error {
	dataset := client.Dataset(configs.DatasetID)
	// Verificar si el dataset ya existe
	if _, err := dataset.Metadata(ctx); err == nil {
		log.Println("Dataset already exists, continuing...")
		return nil
	} else if apiErr, ok := err.(*googleapi.Error); ok && apiErr.Code != 404 {
		return fmt.Errorf("failed to check dataset existence: %v", err)
	}

	// Crear el dataset si no existe
	if err := dataset.Create(ctx, &bigquery.DatasetMetadata{
		Location: "US",
	}); err != nil {
		return fmt.Errorf("failed to create dataset: %v", err)
	}
	log.Println("Dataset created successfully")
	return nil
}

func createStationTable(ctx context.Context, client *bigquery.Client) error {
	table := client.Dataset(configs.DatasetID).Table(configs.TableStations)
	schema := bigquery.Schema{
		{Name: "id", Type: bigquery.StringFieldType},
		{Name: "nombre", Type: bigquery.StringFieldType},
		{Name: "provincia", Type: bigquery.StringFieldType},
		{Name: "latitud", Type: bigquery.StringFieldType},
		{Name: "longitud", Type: bigquery.StringFieldType},
		{Name: "inicio", Type: bigquery.StringFieldType},
		{Name: "final", Type: bigquery.StringFieldType},
		{Name: "faltantes", Type: bigquery.StringFieldType, Repeated: true},
	}
	if err := table.Create(ctx, &bigquery.TableMetadata{
		Schema: schema,
	}); err != nil {
		if apiErr, ok := err.(*googleapi.Error); ok && apiErr.Code == 409 {
			log.Println("Table already exists, continuing...")
			return nil
		}
		return fmt.Errorf("failed to create table: %v", err)
	}
	log.Println("Table created successfully")
	return nil
}
func createWeatherDataTable(ctx context.Context, client *bigquery.Client) error {
	table := client.Dataset(configs.DatasetID).Table(configs.TableWeather)
	schema := bigquery.Schema{
		{Name: "fecha", Type: bigquery.StringFieldType},
		{Name: "indicativo", Type: bigquery.StringFieldType},
		{Name: "nombre", Type: bigquery.StringFieldType},
		{Name: "provincia", Type: bigquery.StringFieldType},
		{Name: "altitud", Type: bigquery.StringFieldType},
		{Name: "tmed", Type: bigquery.StringFieldType},
		{Name: "prec", Type: bigquery.StringFieldType},
		{Name: "tmin", Type: bigquery.StringFieldType},
		{Name: "horatmin", Type: bigquery.StringFieldType},
		{Name: "tmax", Type: bigquery.StringFieldType},
		{Name: "horatmax", Type: bigquery.StringFieldType},
		{Name: "dir", Type: bigquery.StringFieldType},
		{Name: "velmedia", Type: bigquery.StringFieldType},
		{Name: "racha", Type: bigquery.StringFieldType},
		{Name: "horaracha", Type: bigquery.StringFieldType},
		{Name: "presMax", Type: bigquery.StringFieldType},
		{Name: "horaPresMax", Type: bigquery.StringFieldType},
		{Name: "presMin", Type: bigquery.StringFieldType},
		{Name: "horaPresMin", Type: bigquery.StringFieldType},
		{Name: "hrMedia", Type: bigquery.StringFieldType},
		{Name: "hrMax", Type: bigquery.StringFieldType},
		{Name: "horaHrMax", Type: bigquery.StringFieldType},
		{Name: "hrMin", Type: bigquery.StringFieldType},
		{Name: "horaHrMin", Type: bigquery.StringFieldType},
	}

	if err := table.Create(ctx, &bigquery.TableMetadata{
		Schema: schema,
	}); err != nil {
		if apiErr, ok := err.(*googleapi.Error); ok && apiErr.Code == 409 {
			log.Println("Table already exists, continuing...")
			return nil
		}
		return fmt.Errorf("failed to create table: %v", err)
	}
	log.Println("Table created successfully")
	return nil
}

func createErrorTable(ctx context.Context, client *bigquery.Client) error {
	table := client.Dataset(configs.DatasetID).Table(configs.TableErrors)
	schema := bigquery.Schema{
		{Name: "problem_date", Type: bigquery.StringFieldType},
		{Name: "table", Type: bigquery.StringFieldType},
		{Name: "station_id", Type: bigquery.StringFieldType},
		{Name: "message", Type: bigquery.StringFieldType},
		{Name: "timestamp", Type: bigquery.TimestampFieldType},
		{Name: "fixed", Type: bigquery.BooleanFieldType},
	}
	if err := table.Create(ctx, &bigquery.TableMetadata{
		Schema: schema,
	}); err != nil {
		if apiErr, ok := err.(*googleapi.Error); ok && apiErr.Code == 409 {
			log.Println("Table already exists, continuing...")
			return nil
		}
		return fmt.Errorf("failed to create table: %v", err)
	}
	log.Println("Table created successfully")
	return nil
}
