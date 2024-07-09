package configs

// Variable global para la fecha fija
var FixedDate = "2023-01-01T00:00:00UTC"

const (
	// Formato correcto basado en la referencia fija
	DateFormat      = "2006-01-02T15:04:05UTC"
	DateFormatUTC   = "2006-01-02T15:04:05"
	ShortDateFormat = "2006-01-02"
)

// Plantillas de la API de AEMET
var BaseURL = "https://opendata.aemet.es/opendata/api/valores/climatologicos/"
var EstacionesData = BaseURL + "inventarioestaciones/todasestaciones/?api_key=%s"
var EstacionURL = BaseURL + "diarios/datos/fechaini/%s/fechafin/%s/estacion/%s?api_key=%s"

// Constantes para fechas y chunk size
const (
	StartDate = "2024-03-01"
	ChunkSize = 30 * 6
)

// Otras configuraciones globales
var CredPath = "gcloud-key.json"

// Configuraciones para BigQuery
const (
	ProjectID     = "challenge-4-426500"
	DatasetID     = "AEMET_Staging"
	TableStations = "Stations_State"
	TableWeather  = "Weather_entry"
	TableErrors   = "errors"
)
