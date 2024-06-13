package post

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

func AddDistanceToAsteroID(c echo.Context) error {
	ctx, cancel, AsteroidCollection, err := handlers.SetupDBContext()
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}
	defer cancel()

	asteroidID := c.Param("id")
	var objID primitive.ObjectID

	// Intentar convertir a ObjectID para id y nombre
	objID, err = primitive.ObjectIDFromHex(asteroidID)
	if err != nil {
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

	distance := models.Distance{}
	if err := c.Bind(&distance); err != nil {
		return handlers.BindErrorJSON(c, err)
	}

	// Validar los datos de la nueva distancia
	if err := handlers.ValidateDistance(distance); err != nil {
		return handlers.CustomValidationErrorJSON(c, err.Error())
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

	// Verificar si la fecha de la nueva distancia ya existe en las distancias del asteroide
	for _, d := range existingAsteroid.Distances {
		if d.Date == distance.Date {
			return handlers.CustomValidationErrorJSON(c, "Distance for this date already exists")
		}
	}

	// Agregar la nueva distancia al asteroide
	existingAsteroid.Distances = append(existingAsteroid.Distances, distance)

	// Actualizar el asteroide en la base de datos
	_, err = AsteroidCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": existingAsteroid})
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, responses.AsteroidResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &echo.Map{"data": "Distance added to asteroid"},
	})
}
