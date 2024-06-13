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

func CreateAsteroid(c echo.Context) error {
	ctx, cancel, AsteroidCollection, err := handlers.SetupDBContext()
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}
	defer cancel()

	var asteroid models.Asteroid
	if err := c.Bind(&asteroid); err != nil {
		return handlers.BindErrorJSON(c, err)
	}

	// Validar los datos del asteroide
	if err := handlers.ValidateAsteroid(asteroid.Diameter, asteroid.DiscoveryDate, asteroid.Distances); err != nil {
		return handlers.CustomValidationErrorJSON(c, err.Error())
	}

	MappingCollection := configs.GetMappingCollection(configs.ConnectDB(), "asteroid_mappings")
	// Verificar si el nombre del asteroide ya existe
	var existingMapping bson.M
	err = MappingCollection.FindOne(ctx, bson.M{"name": asteroid.Name}).Decode(&existingMapping)
	if err == nil {
		return handlers.CustomValidationErrorJSON(c, "Asteroid with this name already exists")
	} else if err != mongo.ErrNoDocuments {
		return handlers.InternalServerErrorResponse(c, err)
	}

	asteroid.ID = primitive.NewObjectID()

	// Insertar el asteroide en la colección principal
	_, err = AsteroidCollection.InsertOne(ctx, asteroid)
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}

	// Insertar el mapeo en la colección de mapeo
	_, err = MappingCollection.InsertOne(ctx, bson.M{"name": asteroid.Name, "id": asteroid.ID})
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, responses.AsteroidResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    &echo.Map{"data": bson.M{"InsertedID": asteroid.ID}},
	})
}
