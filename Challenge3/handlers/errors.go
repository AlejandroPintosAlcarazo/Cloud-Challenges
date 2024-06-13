package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InternalServerErrorResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"status":  http.StatusInternalServerError,
		"message": "error",
		"data":    map[string]string{"error": err.Error()},
	})
}

func BindErrorJSON(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"status":  http.StatusBadRequest,
		"message": "error",
		"data":    map[string]string{"error": err.Error()},
	})
}

func ValidationErrorJSON(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"status":  http.StatusBadRequest,
		"message": "error",
		"data":    map[string]string{"error": err.Error()},
	})
}

func AsteroidExistsErrorJSON(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"status":  http.StatusBadRequest,
		"message": "error",
		"data":    map[string]string{"data": message},
	})
}

func CustomValidationErrorJSON(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"status":  http.StatusBadRequest,
		"message": "error",
		"data":    map[string]string{"error": message},
	})
}
