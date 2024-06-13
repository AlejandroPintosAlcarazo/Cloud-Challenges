package patch

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

func UpdateAsteroid(c echo.Context) error {
	ctx, cancel, AsteroidCollection, err := handlers.SetupDBContext()
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}
	defer cancel()

	asteroidID := c.Param("id")
	var objID primitive.ObjectID

	// Intentar convertir a ObjectID
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

	var asteroid models.Asteroid
	if err := c.Bind(&asteroid); err != nil {
		return handlers.BindErrorJSON(c, err)
	}

	// Validar los datos del asteroide para la operación PATCH
	if err := handlers.ValidatePatch(&asteroid); err != nil {
		return handlers.CustomValidationErrorJSON(c, err.Error())
	}

	// Preparar los campos actualizados
	updateFields, err := handlers.PrepareUpdateFields(&asteroid)
	if err != nil {
		return handlers.CustomValidationErrorJSON(c, err.Error())
	}

	// Actualizar el asteroide en la base de datos
	update := bson.M{
		"$set": updateFields,
	}

	// Imprimir el ID y la actualización para depuración
	c.Logger().Infof("Updating asteroid with ID: %s with data: %+v", objID.Hex(), update)

	result, err := AsteroidCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}

	// Verificar el resultado de la actualización
	if result.MatchedCount == 0 {
		return handlers.AsteroidExistsErrorJSON(c, "Asteroid does not exist")
	}

	// Actualizar el mapeo en la colección de mapeo
	if _, ok := updateFields["name"]; ok {
		MappingCollection := configs.GetMappingCollection(configs.ConnectDB(), "asteroid_mappings")
		_, err = MappingCollection.UpdateOne(ctx, bson.M{"id": objID}, bson.M{"$set": bson.M{"name": asteroid.Name}})
		if err != nil {
			return handlers.InternalServerErrorResponse(c, err)
		}
	}

	return c.JSON(http.StatusOK, responses.AsteroidResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &echo.Map{"data": "Asteroid updated successfully"},
	})
}
