package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {

	r := gin.Default()
	r.GET("/v1/health", app.healthHandler)
	r.GET("/v1//quotas/:id", app.showQuotaHandler)
	r.POST("/v1//quotas", app.createQuotaHandler)

	return r
}
