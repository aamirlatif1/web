package main

import (
	"net/http"
	"time"

	"github.com/aamirlatif1/web/internal/data"
	"github.com/gin-gonic/gin"
)

type envolope map[string]any

func (app *application) showQuotaHandler(c *gin.Context) {
	id, err := app.readIDParam(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "404 page not found"})
		return
	}

	quota := data.Quota{
		QuotaID:      id,
		ResourceType: "Storage",
		Scope:        "test",
		ScopeID:      "scope1",
		Limits: data.Limits{
			HardLimit:       100,
			SoftLimit:       80,
			RefreshInterval: "monthly",
		},
		Usage: data.Usage{
			Current:     45,
			LastRefresh: time.Now(),
			// NextRefresh: time.Now(),
		},
		Status: data.StatusOK,
	}

	c.JSON(http.StatusOK, envolope{"quota": quota})
}

func (app *application) createQuotaHandler(c *gin.Context) {
}
