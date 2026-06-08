package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/services"
	"github.com/pablomillaquen/speckit_golang_api/pkg/response"
)

type MaterialHandler struct {
	service *services.MaterialService
}

func NewMaterialHandler(service *services.MaterialService) *MaterialHandler {
	return &MaterialHandler{service: service}
}

// @Summary Create material
// @Description Create a new material in the inventory catalog
// @Tags Materials
// @Accept json
// @Produce json
// @Param request body dto.CreateMaterialRequest true "Material data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/materials [post]
func (h *MaterialHandler) Create(c *gin.Context) {
	var req dto.CreateMaterialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.service.Create(req)
	if err != nil {
		response.Conflict(c, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Material created", resp)
}

// @Summary Get material by ID
// @Description Retrieve a single material by ID
// @Tags Materials
// @Produce json
// @Param id path int true "Material ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/v1/materials/{id} [get]
func (h *MaterialHandler) FindByID(c *gin.Context) {
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
	response.Success(c, http.StatusOK, "Material found", resp)
}

// @Summary List materials
// @Description Retrieve a paginated list of materials
// @Tags Materials
// @Produce json
// @Param page query int false "Page number"
// @Param per_page query int false "Items per page"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/materials [get]
func (h *MaterialHandler) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}
	materials, total, err := h.service.FindAll(page, perPage)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.SuccessPaginated(c, "Materials retrieved", materials, page, perPage, total)
}

// @Summary Update material
// @Description Update an existing material
// @Tags Materials
// @Accept json
// @Produce json
// @Param id path int true "Material ID"
// @Param request body dto.UpdateMaterialRequest true "Updated material data"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/materials/{id} [put]
func (h *MaterialHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	var req dto.UpdateMaterialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.service.Update(id, req)
	if err != nil {
		response.Conflict(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Material updated", resp)
}

// @Summary Delete material
// @Description Delete a material by ID
// @Tags Materials
// @Produce json
// @Param id path int true "Material ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/materials/{id} [delete]
func (h *MaterialHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	if err := h.service.Delete(id); err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Material deleted", nil)
}
