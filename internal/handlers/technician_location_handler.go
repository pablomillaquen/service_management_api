package handlers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/services"
	"github.com/pablomillaquen/speckit_golang_api/pkg/response"
)

type TechnicianLocationHandler struct {
	service *services.TechnicianLocationService
}

func NewTechnicianLocationHandler(service *services.TechnicianLocationService) *TechnicianLocationHandler {
	return &TechnicianLocationHandler{service: service}
}

// @Summary Report technician location
// @Description Report current GPS location (technician only). The user_id is automatically set from the authenticated token.
// @Tags Technician Locations
// @Accept json
// @Produce json
// @Param request body dto.CreateTechnicianLocationRequest true "GPS Coordinates"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Router /api/v1/technician-locations [post]
func (h *TechnicianLocationHandler) Create(c *gin.Context) {
	var req dto.CreateTechnicianLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	userID, _ := c.Get("user_id")
	uid, _ := userID.(float64)
	resp, err := h.service.Create(uint64(uid), req)
	if err != nil {
		if err == services.ErrInvalidLatitude || err == services.ErrInvalidLongitude {
			response.ValidationError(c, err.Error(), nil)
			return
		}
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, 201, "Location reported successfully", resp)
}

// @Summary List technician locations
// @Description Query technician locations with optional filters. Administrators see all; technicians see only their own.
// @Tags Technician Locations
// @Produce json
// @Param user_id query int false "Filter by technician ID (admin only)"
// @Param start_date query string false "ISO 8601 start datetime filter (inclusive)"
// @Param end_date query string false "ISO 8601 end datetime filter (inclusive)"
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 20)"
// @Success 200 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Failure 403 {object} response.APIResponse
// @Router /api/v1/technician-locations [get]
func (h *TechnicianLocationHandler) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 20
	}
	userRole, _ := c.Get("user_role")
	role, _ := userRole.(string)
	var userIDFilter *uint64
	if role == "technician" {
		uid, _ := c.Get("user_id")
		uidVal, _ := uid.(float64)
		u := uint64(uidVal)
		userIDFilter = &u
	}
	var queryUserID *uint64
	if qUID := c.Query("user_id"); qUID != "" {
		if id, err := strconv.ParseUint(qUID, 10, 64); err == nil {
			queryUserID = &id
		}
	}
	var startDate, endDate *time.Time
	if sd := c.Query("start_date"); sd != "" {
		if t, err := time.Parse(time.RFC3339, sd); err == nil {
			startDate = &t
		}
	}
	if ed := c.Query("end_date"); ed != "" {
		if t, err := time.Parse(time.RFC3339, ed); err == nil {
			endDate = &t
		}
	}
	effectiveUserID := userIDFilter
	if effectiveUserID == nil {
		effectiveUserID = queryUserID
	}
	locations, total, err := h.service.FindAll(page, perPage, effectiveUserID, startDate, endDate)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.SuccessPaginated(c, "Locations retrieved", locations, page, perPage, total)
}
