package fetcher

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/models"
)

// processInitialResponse processes the initial response from AEMET API
func parseInitialResponse(respBody io.ReadCloser) (models.InitialResponse, error) {
	// Read the response body
	body, err := io.ReadAll(respBody)
	if err != nil {
		return models.InitialResponse{}, err
	}

	// Print the response body for debugging
	//fmt.Println("Initial Response Body:", string(body))

	// Parse the JSON response
	var initialResponse models.InitialResponse
	if err := json.Unmarshal(body, &initialResponse); err != nil {
		return models.InitialResponse{}, err
	}

	// Check if the "datos" field is present
	if initialResponse.Datos == "" {
		return models.InitialResponse{}, fmt.Errorf("error: no data URL found in response")
	}

	return initialResponse, nil
}

// parseResponseBody parses the JSON response body into AEMETResponse
func parseActualDataResponse(dataBody []byte) (*models.AEMETResponse, error) {
	// Parse the JSON response
	var data models.AEMETResponse
	if err := json.Unmarshal(dataBody, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// parseGeoStations analiza y convierte los datos de estaciones geográficas
func parseStations(dataBody []byte) ([]models.SimpleStation, error) {
	// Analiza la respuesta de los datos reales en una estructura genérica
	var result []map[string]interface{}
	if err := json.Unmarshal(dataBody, &result); err != nil {
		return nil, fmt.Errorf("error unmarshalling actual data: %v", err)
	}

	// Convierte la estructura genérica en una lista de GeoEstacion
	estaciones := make([]models.SimpleStation, len(result))
	for i, item := range result {
		indicativo, _ := item["indicativo"].(string)
		nombre, _ := item["nombre"].(string)
		provincia, _ := item["provincia"].(string)
		latitud, _ := item["latitud"].(string)
		longitud, _ := item["longitud"].(string)

		estaciones[i] = models.SimpleStation{
			ID:        indicativo,
			Nombre:    nombre,
			Provincia: provincia,
			Latitud:   latitud,
			Longitud:  longitud,
		}
	}

	return estaciones, nil
}
