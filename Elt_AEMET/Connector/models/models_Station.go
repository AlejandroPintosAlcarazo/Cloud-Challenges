package models

// Station represents the structure for the stations
type Station struct {
	ID        string   `json:"id" bigquery:"id"`
	Nombre    string   `json:"nombre" bigquery:"nombre"`
	Provincia string   `json:"provincia" bigquery:"provincia"`
	Latitud   string   `json:"latitud" bigquery:"latitud"`
	Longitud  string   `json:"longitud" bigquery:"longitud"`
	Inicio    string   `json:"inicio" bigquery:"inicio"`
	Final     string   `json:"final" bigquery:"final"`
	Faltantes []string `json:"faltantes" bigquery:"faltantes"`
}

type SimpleStation struct {
	ID        string `json:"id"`
	Nombre    string `json:"nombre"`
	Provincia string `json:"provincia"`
	Latitud   string `json:"latitud"`
	Longitud  string `json:"longitud"`
}

type StationState struct {
	ID        string   `bigquery:"id"`
	Final     string   `bigquery:"final"`
	Faltantes []string `json:"faltantes" bigquery:"faltantes"`
}

type StationUpdate struct {
	ID    string `bigquery:"id"`
	Final string `bigquery:"final"`
}
