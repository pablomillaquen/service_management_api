package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pablomillaquen/speckit_golang_api/internal/services"
	"github.com/pablomillaquen/speckit_golang_api/pkg/response"
)

type AuditHandler struct {
	service *services.AuditService
}

func NewAuditHandler(service *services.AuditService) *AuditHandler {
	return &AuditHandler{service: service}
}

// @Summary List audit logs
// @Description Retrieve a paginated list of audit logs with optional filters
// @Tags Audit
// @Produce json
// @Param page query int false "Page number"
// @Param per_page query int false "Items per page"
// @Param entity query string false "Filter by entity name"
// @Param entity_id query int false "Filter by entity ID"
// @Param user_id query int false "Filter by user ID"
// @Param action query string false "Filter by action"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/audit [get]
func (h *AuditHandler) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}
	filters := map[string]interface{}{
		"entity":    c.Query("entity"),
		"entity_id": c.Query("entity_id"),
		"user_id":   c.Query("user_id"),
		"action":    c.Query("action"),
	}
	logs, total, err := h.service.FindAll(page, perPage, filters)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.SuccessPaginated(c, "Audit logs retrieved", logs, page, perPage, total)
}
