package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/services"
	"github.com/pablomillaquen/speckit_golang_api/pkg/response"
)

type CatalogHandler struct {
	typeService  *services.EquipmentTypeService
	brandService *services.BrandService
	modelService *services.EquipmentModelService
}

func NewCatalogHandler(
	typeService *services.EquipmentTypeService,
	brandService *services.BrandService,
	modelService *services.EquipmentModelService,
) *CatalogHandler {
	return &CatalogHandler{
		typeService:  typeService,
		brandService: brandService,
		modelService: modelService,
	}
}

// @Summary Create equipment type
// @Description Create a new equipment type in the catalog
// @Tags Catalog
// @Accept json
// @Produce json
// @Param request body dto.CreateEquipmentTypeRequest true "Equipment type data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/catalog/types [post]
func (h *CatalogHandler) CreateType(c *gin.Context) {
	var req dto.CreateEquipmentTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.typeService.Create(req)
	if err != nil {
		response.Conflict(c, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Equipment type created", resp)
}

// @Summary Get equipment type
// @Description Retrieve a single equipment type by ID
// @Tags Catalog
// @Produce json
// @Param id path int true "Equipment type ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/v1/catalog/types/{id} [get]
func (h *CatalogHandler) GetType(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	resp, err := h.typeService.FindByID(id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Equipment type found", resp)
}

// @Summary List equipment types
// @Description Retrieve all equipment types
// @Tags Catalog
// @Produce json
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/catalog/types [get]
func (h *CatalogHandler) ListTypes(c *gin.Context) {
	types, err := h.typeService.FindAll()
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Equipment types retrieved", types)
}

// @Summary Update equipment type
// @Description Update an existing equipment type
// @Tags Catalog
// @Accept json
// @Produce json
// @Param id path int true "Equipment type ID"
// @Param request body dto.UpdateEquipmentTypeRequest true "Updated equipment type data"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/catalog/types/{id} [put]
func (h *CatalogHandler) UpdateType(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	var req dto.UpdateEquipmentTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.typeService.Update(id, req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusOK, "Equipment type updated", resp)
}

// @Summary Delete equipment type
// @Description Delete an equipment type by ID
// @Tags Catalog
// @Produce json
// @Param id path int true "Equipment type ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/catalog/types/{id} [delete]
func (h *CatalogHandler) DeleteType(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	if err := h.typeService.Delete(id); err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Equipment type deleted", nil)
}

// @Summary Create brand
// @Description Create a new brand in the catalog
// @Tags Catalog
// @Accept json
// @Produce json
// @Param request body dto.CreateBrandRequest true "Brand data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/catalog/brands [post]
func (h *CatalogHandler) CreateBrand(c *gin.Context) {
	var req dto.CreateBrandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.brandService.Create(req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusCreated, "Brand created", resp)
}

// @Summary Get brand by ID
// @Description Retrieve a single brand by ID
// @Tags Catalog
// @Produce json
// @Param id path int true "Brand ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/v1/catalog/brands/{id} [get]
func (h *CatalogHandler) GetBrand(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	resp, err := h.brandService.FindByID(id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Brand found", resp)
}

// @Summary List brands by type
// @Description Retrieve all brands for a given equipment type
// @Tags Catalog
// @Produce json
// @Param typeId path int true "Equipment type ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/catalog/brands/by-type/{typeId} [get]
func (h *CatalogHandler) ListBrandsByType(c *gin.Context) {
	typeID, err := strconv.ParseUint(c.Param("typeId"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid type ID", nil)
		return
	}
	brands, err := h.brandService.FindByTypeID(typeID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Brands retrieved", brands)
}

// @Summary Update brand
// @Description Update an existing brand
// @Tags Catalog
// @Accept json
// @Produce json
// @Param id path int true "Brand ID"
// @Param request body dto.UpdateBrandRequest true "Updated brand data"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/catalog/brands/{id} [put]
func (h *CatalogHandler) UpdateBrand(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	var req dto.UpdateBrandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.brandService.Update(id, req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusOK, "Brand updated", resp)
}

// @Summary Delete brand
// @Description Delete a brand by ID
// @Tags Catalog
// @Produce json
// @Param id path int true "Brand ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/catalog/brands/{id} [delete]
func (h *CatalogHandler) DeleteBrand(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	if err := h.brandService.Delete(id); err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Brand deleted", nil)
}

// @Summary Create equipment model
// @Description Create a new equipment model in the catalog
// @Tags Catalog
// @Accept json
// @Produce json
// @Param request body dto.CreateEquipmentModelRequest true "Equipment model data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/catalog/models [post]
func (h *CatalogHandler) CreateModel(c *gin.Context) {
	var req dto.CreateEquipmentModelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.modelService.Create(req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusCreated, "Equipment model created", resp)
}

// @Summary Get equipment model
// @Description Retrieve a single equipment model by ID
// @Tags Catalog
// @Produce json
// @Param id path int true "Equipment model ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/v1/catalog/models/{id} [get]
func (h *CatalogHandler) GetModel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	resp, err := h.modelService.FindByID(id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Equipment model found", resp)
}

// @Summary List models by brand
// @Description Retrieve all equipment models for a given brand
// @Tags Catalog
// @Produce json
// @Param brandId path int true "Brand ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/catalog/models/by-brand/{brandId} [get]
func (h *CatalogHandler) ListModelsByBrand(c *gin.Context) {
	brandID, err := strconv.ParseUint(c.Param("brandId"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid brand ID", nil)
		return
	}
	models, err := h.modelService.FindByBrandID(brandID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Equipment models retrieved", models)
}

// @Summary Update equipment model
// @Description Update an existing equipment model
// @Tags Catalog
// @Accept json
// @Produce json
// @Param id path int true "Equipment model ID"
// @Param request body dto.UpdateEquipmentModelRequest true "Updated equipment model data"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/catalog/models/{id} [put]
func (h *CatalogHandler) UpdateModel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	var req dto.UpdateEquipmentModelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.modelService.Update(id, req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusOK, "Equipment model updated", resp)
}

// @Summary Delete equipment model
// @Description Delete an equipment model by ID
// @Tags Catalog
// @Produce json
// @Param id path int true "Equipment model ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/catalog/models/{id} [delete]
func (h *CatalogHandler) DeleteModel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	if err := h.modelService.Delete(id); err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Equipment model deleted", nil)
}
