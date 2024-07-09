package models

// InitialResponse represents the structure of the initial response from AEMET API
type InitialResponse struct {
	Descripcion string `json:"descripcion"`
	Estado      int    `json:"estado"`
	Datos       string `json:"datos"`
	Metadatos   string `json:"metadatos"`
}

// DataEntry represents a single entry in the final data response from AEMET API
type DataEntry struct {
	Date          string `json:"fecha" bigquery:"fecha"`
	Indicativo    string `json:"indicativo" bigquery:"indicativo"`
	Nombre        string `json:"nombre" bigquery:"nombre"`
	Provincia     string `json:"provincia" bigquery:"provincia"`
	Altitud       string `json:"altitud" bigquery:"altitud"`
	Temperature   string `json:"tmed" bigquery:"tmed"`
	Precipitation string `json:"prec" bigquery:"prec"`
	Tmin          string `json:"tmin" bigquery:"tmin"`
	HoraTmin      string `json:"horatmin" bigquery:"horatmin"`
	Tmax          string `json:"tmax" bigquery:"tmax"`
	HoraTmax      string `json:"horatmax" bigquery:"horatmax"`
	Dir           string `json:"dir" bigquery:"dir"`
	Velmedia      string `json:"velmedia" bigquery:"velmedia"`
	Racha         string `json:"racha" bigquery:"racha"`
	HoraRacha     string `json:"horaracha" bigquery:"horaracha"`
	PresMax       string `json:"presMax" bigquery:"presMax"`
	HoraPresMax   string `json:"horaPresMax" bigquery:"horaPresMax"`
	PresMin       string `json:"presMin" bigquery:"presMin"`
	HoraPresMin   string `json:"horaPresMin" bigquery:"horaPresMin"`
	HrMedia       string `json:"hrMedia" bigquery:"hrMedia"`
	HrMax         string `json:"hrMax" bigquery:"hrMax"`
	HoraHrMax     string `json:"horaHrMax" bigquery:"horaHrMax"`
	HrMin         string `json:"hrMin" bigquery:"hrMin"`
	HoraHrMin     string `json:"horaHrMin" bigquery:"horaHrMin"`
}

// AEMETResponse represents the structure of the final data response from AEMET API
type AEMETResponse []DataEntry
