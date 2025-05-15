package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	})
}
