package main

import (
	"os"

	"github.com/AlejandroPintosAlcarazo/asteroid.API/configs"
	"github.com/AlejandroPintosAlcarazo/asteroid.API/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configs.ConnectDB()
	routes.AsteroidRoute(e)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "hola mongo y go"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
