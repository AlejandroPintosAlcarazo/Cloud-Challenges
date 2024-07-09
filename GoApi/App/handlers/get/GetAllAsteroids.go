package get

import (
	"net/http"
	"strconv"

	"github.com/AlejandroPintosAlcarazo/asteroid.API/handlers"
	"github.com/AlejandroPintosAlcarazo/asteroid.API/models"
	"github.com/AlejandroPintosAlcarazo/asteroid.API/responses"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllAsteroids(c echo.Context) error {
	ctx, cancel, AsteroidCollection, err := handlers.SetupDBContext()
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}
	defer cancel()

	// Get pagination parameters from query
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	// Convert parameters to integers with default values
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	// Calculate the number of documents to skip
	skip := (page - 1) * limit

	// Get the total count of documents
	totalCount, err := AsteroidCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}

	// Calculate the total number of pages
	totalPages := int((totalCount + int64(limit) - 1) / int64(limit)) // Redondeo hacia arriba

	// If the requested page is greater than the total pages, return an error
	if page > totalPages {
		return c.JSON(http.StatusBadRequest, responses.AsteroidResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid page number",
			Data:    &echo.Map{"data": "Requested page exceeds total number of pages"},
		})
	}

	// Find options with limit and skip
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(skip))

	// Execute the query with pagination
	cursor, err := AsteroidCollection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}
	defer cursor.Close(ctx)

	var asteroids []models.Asteroid
	for cursor.Next(ctx) {
		var asteroid models.Asteroid
		if err := cursor.Decode(&asteroid); err != nil {
			return handlers.InternalServerErrorResponse(c, err)
		}
		asteroids = append(asteroids, asteroid)
	}

	if err := cursor.Err(); err != nil {
		return handlers.InternalServerErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, responses.AsteroidResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &echo.Map{"data": asteroids},
	})
}
