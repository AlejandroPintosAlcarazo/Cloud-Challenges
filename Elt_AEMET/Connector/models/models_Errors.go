package models

import "time"

type ServerErrorRecord struct {
	ProblemDate string    `bigquery:"problem_date"`
	Table       string    `bigquery:"table"`
	StationID   string    `bigquery:"station_id"`
	Message     string    `bigquery:"message"`
	Timestamp   time.Time `bigquery:"timestamp"`
	Fixed       bool      `bigquery:"fixed"`
}
