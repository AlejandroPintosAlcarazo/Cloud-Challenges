package delete

import (
	"net/http"

	"github.com/AlejandroPintosAlcarazo/asteroid.API/configs"
	"github.com/AlejandroPintosAlcarazo/asteroid.API/handlers"
	"github.com/AlejandroPintosAlcarazo/asteroid.API/models"
	"github.com/AlejandroPintosAlcarazo/asteroid.API/responses"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteDistance(c echo.Context) error {
	ctx, cancel, AsteroidCollection, err := handlers.SetupDBContext()
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}
	defer cancel()

	asteroidID := c.Param("id")
	distanceID := c.Param("distanceID")
	var objID primitive.ObjectID

	// Intentar convertir a ObjectID
	objID, err = primitive.ObjectIDFromHex(asteroidID)
	if err != nil {
		// Si falla, buscar en la colección de mapeo
		MappingCollection := configs.GetMappingCollection(configs.ConnectDB(), "asteroid_mappings")
		var mapping bson.M
		err := MappingCollection.FindOne(ctx, bson.M{"name": asteroidID}).Decode(&mapping)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return handlers.AsteroidExistsErrorJSON(c, "Asteroid does not exist")
			}
			return handlers.InternalServerErrorResponse(c, err)
		}
		objID = mapping["id"].(primitive.ObjectID)
	}

	// Verificar si el asteroide existe en la base de datos
	var existingAsteroid models.Asteroid
	err = AsteroidCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&existingAsteroid)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return handlers.AsteroidExistsErrorJSON(c, "Asteroid does not exist")
		}
		return handlers.InternalServerErrorResponse(c, err)
	}

	// Verificar si la distancia existe
	distanceExists := false
	newDistances := []models.Distance{}
	for _, distance := range existingAsteroid.Distances {
		if distance.Date == distanceID {
			distanceExists = true
		} else {
			newDistances = append(newDistances, distance)
		}
	}

	if !distanceExists {
		return c.JSON(http.StatusNotFound, responses.AsteroidResponse{
			Status:  http.StatusNotFound,
			Message: "error",
			Data:    &echo.Map{"error": "Distance not found"},
		})
	}

	existingAsteroid.Distances = newDistances

	// Actualizar el asteroide en la base de datos
	_, err = AsteroidCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": existingAsteroid})
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, responses.AsteroidResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &echo.Map{"data": "Distance deleted successfully"},
	})
}
