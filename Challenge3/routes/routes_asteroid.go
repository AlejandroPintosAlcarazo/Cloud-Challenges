package routes

import (
	handlers_delete "github.com/AlejandroPintosAlcarazo/asteroid.API/handlers/delete"
	handlers_get "github.com/AlejandroPintosAlcarazo/asteroid.API/handlers/get"
	handlers_patch "github.com/AlejandroPintosAlcarazo/asteroid.API/handlers/patch"
	handlers_post "github.com/AlejandroPintosAlcarazo/asteroid.API/handlers/post"
	"github.com/labstack/echo/v4"
)

func AsteroidRoute(e *echo.Echo) {
	api := e.Group("/api/v1", serverHeader)

	api.DELETE("/asteroids", handlers_delete.DeleteAllAsteroids)
	api.DELETE("/asteroids/:id", handlers_delete.DeleteAsteroidByID)
	api.DELETE("/asteroids/:id/distances/:distanceID", handlers_delete.DeleteDistance)

	api.GET("/asteroids", handlers_get.GetAllAsteroids)
	api.GET("/asteroids/:id", handlers_get.GetAsteroidByID)

	api.PATCH("/asteroids/:id", handlers_patch.UpdateAsteroid)

	api.POST("/asteroids/:id/distances", handlers_post.AddDistanceToAsteroID)
	api.POST("/asteroids", handlers_post.CreateAsteroid)

}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("x-version", "Test/v1.0")
		return next(c)
	}
}
