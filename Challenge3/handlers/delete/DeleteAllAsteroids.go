package delete

import (
	"net/http"

	"github.com/AlejandroPintosAlcarazo/asteroid.API/configs"
	"github.com/AlejandroPintosAlcarazo/asteroid.API/handlers"
	"github.com/AlejandroPintosAlcarazo/asteroid.API/responses"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteAllAsteroids(c echo.Context) error {
	ctx, cancel, AsteroidCollection, err := handlers.SetupDBContext()
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}
	defer cancel()

	// Eliminar todos los asteroides de la colección principal
	deleteResult, err := AsteroidCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}

	// Limpiar la colección de mapeo
	MappingCollection := configs.GetMappingCollection(configs.ConnectDB(), "asteroid_mappings")
	mappingDeleteResult, err := MappingCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}

	// Verificar si se eliminaron documentos
	if deleteResult.DeletedCount == 0 && mappingDeleteResult.DeletedCount == 0 {
		return c.JSON(http.StatusOK, responses.AsteroidResponse{
			Status:  http.StatusOK,
			Message: "database is already empty",
			Data:    &echo.Map{"data": "No asteroids or mappings to delete"},
		})
	}

	return c.JSON(http.StatusOK, responses.AsteroidResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &echo.Map{"data": "All asteroids and mappings deleted successfully"},
	})
}
