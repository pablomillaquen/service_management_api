package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pablomillaquen/speckit_golang_api/pkg/database"
	"github.com/pablomillaquen/speckit_golang_api/pkg/response"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

// @Summary Health check
// @Description Check if the API service is healthy and database is connected
// @Tags Health
// @Produce json
// @Success 200 {object} response.APIResponse
// @Failure 503 {object} response.APIResponse
// @Router /api/v1/health [get]
func (h *HealthHandler) Check(c *gin.Context) {
	dbReady := database.IsReady(h.db)
	if !dbReady {
		response.Error(c, http.StatusServiceUnavailable, "Service unhealthy", []string{"database not ready"})
		return
	}
	response.Success(c, http.StatusOK, "Service is healthy", gin.H{
		"database": "connected",
	})
}
