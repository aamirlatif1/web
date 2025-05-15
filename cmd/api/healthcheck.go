package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthHandler(c *gin.Context) {
	data := envolope{
		"status": "available",
		"system_info": map[string]string{
			"status":      "available",
			"environment": app.config.env,
			"version":     version,
		},
	}

	c.JSON(http.StatusOK, data)
}
