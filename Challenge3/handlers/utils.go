package handlers

import (
	"context"
	"errors"
	"time"

	"github.com/AlejandroPintosAlcarazo/asteroid.API/configs"
	"github.com/AlejandroPintosAlcarazo/asteroid.API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupDBContext() (context.Context, context.CancelFunc, *mongo.Collection, error) {
	client := configs.ConnectDB()
	asteroidCollection := configs.GetCollection(client, "asteroids")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel, asteroidCollection, nil
}

func ValidateDistance(distance models.Distance) error {
	if distance.Distance <= 0 {
		return errors.New("distance must be greater than 0")
	}

	_, err := time.Parse("2006-01-02", distance.Date)
	if err != nil {
		return errors.New("invalid date format, expected YYYY-MM-DD")
	}

	return nil
}

func ValidateAsteroid(diameter int, discoveryDate string, distances []models.Distance) error {
	if diameter <= 0 {
		return errors.New("diameter must be greater than 0")
	}

	_, err := time.Parse("2006-01-02", discoveryDate)
	if err != nil {
		return errors.New("invalid date format, expected YYYY-MM-DD")
	}

	for _, distance := range distances {
		if err := ValidateDistance(distance); err != nil {
			return err
		}
	}

	return nil
}

func ValidatePatch(asteroid *models.Asteroid) error {

	if asteroid.Diameter != 0 {
		if asteroid.Diameter <= 0 {
			return errors.New("diameter must be greater than 0")
		}
	}

	if asteroid.DiscoveryDate != "" {
		_, err := time.Parse("2006-01-02", asteroid.DiscoveryDate)
		if err != nil {
			return errors.New("invalid date format, expected YYYY-MM-DD")
		}
	}

	if len(asteroid.Distances) > 0 {
		for _, distance := range asteroid.Distances {
			if err := ValidateDistance(distance); err != nil {
				return err
			}
		}
	}

	return nil
}

func PrepareUpdateFields(asteroid *models.Asteroid) (bson.M, error) {
	updateFields := bson.M{}

	if asteroid.Name != "" {
		updateFields["name"] = asteroid.Name
	}

	if asteroid.Diameter != 0 {
		updateFields["diameter"] = asteroid.Diameter
	}

	if asteroid.DiscoveryDate != "" {
		updateFields["discovery_date"] = asteroid.DiscoveryDate
	}

	if asteroid.Observations != "" {
		updateFields["observations"] = asteroid.Observations
	}

	if len(asteroid.Distances) > 0 {
		updateFields["distances"] = asteroid.Distances
	}

	if len(updateFields) == 0 {
		return nil, errors.New("no valid fields to update")
	}

	return updateFields, nil
}
