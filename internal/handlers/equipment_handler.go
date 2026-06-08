package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/services"
	"github.com/pablomillaquen/speckit_golang_api/pkg/response"
)

type EquipmentHandler struct {
	service *services.EquipmentService
}

func NewEquipmentHandler(service *services.EquipmentService) *EquipmentHandler {
	return &EquipmentHandler{service: service}
}

// @Summary Create equipment
// @Description Register new equipment for a client
// @Tags Equipment
// @Accept json
// @Produce json
// @Param request body dto.CreateEquipmentRequest true "Equipment data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/equipment [post]
func (h *EquipmentHandler) Create(c *gin.Context) {
	var req dto.CreateEquipmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.service.Create(req)
	if err != nil {
		response.Conflict(c, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Equipment created", resp)
}

// @Summary Get equipment by ID
// @Description Retrieve a single piece of equipment by ID
// @Tags Equipment
// @Produce json
// @Param id path int true "Equipment ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/v1/equipment/{id} [get]
func (h *EquipmentHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	resp, err := h.service.FindByID(id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Equipment found", resp)
}

// @Summary List equipment
// @Description Retrieve a paginated list of equipment with optional filters
// @Tags Equipment
// @Produce json
// @Param page query int false "Page number"
// @Param per_page query int false "Items per page"
// @Param search query string false "Search term"
// @Param client_id query int false "Filter by client ID"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/equipment [get]
func (h *EquipmentHandler) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}
	search := c.Query("search")
	clientID, _ := strconv.ParseUint(c.Query("client_id"), 10, 64)
	equipments, total, err := h.service.FindAll(page, perPage, search, clientID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.SuccessPaginated(c, "Equipment retrieved", equipments, page, perPage, total)
}

// @Summary Update equipment
// @Description Update an existing piece of equipment
// @Tags Equipment
// @Accept json
// @Produce json
// @Param id path int true "Equipment ID"
// @Param request body dto.UpdateEquipmentRequest true "Updated equipment data"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/equipment/{id} [put]
func (h *EquipmentHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	var req dto.UpdateEquipmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.service.Update(id, req)
	if err != nil {
		response.Conflict(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Equipment updated", resp)
}

// @Summary Delete equipment
// @Description Delete a piece of equipment by ID
// @Tags Equipment
// @Produce json
// @Param id path int true "Equipment ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/equipment/{id} [delete]
func (h *EquipmentHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	if err := h.service.Delete(id); err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Equipment deleted", nil)
}
