package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func fetchDataHandler(c *gin.Context) {
	data, err := fetchSatelliteData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = saveToDatabase(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data fetched and saved successfully"})
}

func setupEndpoints() {
	r := gin.Default()
	r.GET("/fetch", fetchDataHandler)
	r.Run(":8080")
}

