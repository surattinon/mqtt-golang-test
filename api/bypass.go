package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/conxtech/jsonhandler"
)

var Sensors jsonhandler.SensorData

func getSensors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Sensors)
}

func ApiService() {
	router := gin.Default()
	router.GET("/api/sensors", getSensors)

	router.Run("localhost:8080")
}
