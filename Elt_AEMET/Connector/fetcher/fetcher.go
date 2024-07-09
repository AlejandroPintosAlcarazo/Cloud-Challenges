package fetcher

import (
	"fmt"

	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/configs"
	"github.com/AlejandroPintosAlcarazo/Aemet-ELT/models"
)

// Coordina el proceso de obtención de datos climatológicos de la API de AEMET
func FetchAEMETData(startDate, endDate, stationID string) (*models.AEMETResponse, error) {
	apiKey := configs.LoadApiKey()

	// Transformar las fechas de "yyyy-mm-dd" a "yyyy-mm-ddTHH:MM:SSUTC"
	startDateUTC := startDate + "T00:00:00UTC"
	endDateUTC := endDate + "T23:59:59UTC"

	// Construir el endpoint usando variables globales
	endpoint := fmt.Sprintf(configs.EstacionURL, startDateUTC, endDateUTC, stationID, apiKey) // nueva línea

	// Realiza la solicitud inicial para obtener la URL de los datos
	resp, err := getInitialDataURL(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Analiza la respuesta inicial para obtener la URL de los datos reales
	initialResponse, err := parseInitialResponse(resp.Body)
	if err != nil {
		return nil, err
	}

	// Realiza la solicitud para obtener los datos reales
	dataBody, err := getActualData(initialResponse.Datos)
	if err != nil {
		return nil, err
	}

	// Analiza la respuesta de los datos reales
	data, err := parseActualDataResponse(dataBody)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func FetchStationsData() ([]models.SimpleStation, error) {
	apiKey := configs.LoadApiKey()

	// Construir el endpoint usando variables globales
	endpoint := fmt.Sprintf(configs.EstacionesData, apiKey)

	// Usa getInitialDataURL para la primera solicitud HTTP
	resp, err := getInitialDataURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error fetching initial data: %v", err)
	}
	defer resp.Body.Close()

	// Usa parseInitialResponse para procesar la respuesta inicial
	initialResponse, err := parseInitialResponse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing initial response: %v", err)
	}

	// Usa getActualData para la segunda solicitud HTTP
	dataBody, err := getActualData(initialResponse.Datos)
	if err != nil {
		return nil, fmt.Errorf("error fetching actual data: %v", err)
	}

	// Usa la función separada para analizar y convertir los datos
	estaciones, err := parseStations(dataBody)
	if err != nil {
		return nil, fmt.Errorf("error parsing and converting data: %v", err)
	}

	return estaciones, nil
}
