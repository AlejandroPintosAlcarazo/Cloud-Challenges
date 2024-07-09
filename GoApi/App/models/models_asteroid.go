package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Distance struct {
	Date     string  `json:"date"`
	Distance float64 `json:"distance"`
}

type Asteroid struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name          string             `json:"name,omitempty" validate:"required"`
	Diameter      int                `json:"diameter,omitempty" validate:"required"`
	DiscoveryDate string             `json:"discovery_date,omitempty" validate:"required"`
	Observations  string             `json:"observations,omitempty"`
	Distances     []Distance         `json:"distances,omitempty"`
}
