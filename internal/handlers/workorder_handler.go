package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/workorder"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/services"
	"github.com/pablomillaquen/speckit_golang_api/pkg/response"
)

type WorkOrderHandler struct {
	service *services.WorkOrderService
}

func NewWorkOrderHandler(service *services.WorkOrderService) *WorkOrderHandler {
	return &WorkOrderHandler{service: service}
}

func getCurrentUserID(c *gin.Context) uint64 {
	userID, _ := c.Get("user_id")
	uid, _ := userID.(uint64)
	return uid
}

// @Summary Create work order
// @Description Create a new work order for a client's equipment
// @Tags Work Orders
// @Accept json
// @Produce json
// @Param request body dto.CreateWorkOrderRequest true "Work order data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /api/v1/work-orders [post]
func (h *WorkOrderHandler) Create(c *gin.Context) {
	userID := getCurrentUserID(c)
	var req dto.CreateWorkOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.service.Create(req, userID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusCreated, "Work order created", resp)
}

// @Summary Get work order by ID
// @Description Retrieve a single work order by its ID
// @Tags Work Orders
// @Produce json
// @Param id path int true "Work order ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/v1/work-orders/{id} [get]
func (h *WorkOrderHandler) FindByID(c *gin.Context) {
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
	response.Success(c, http.StatusOK, "Work order found", resp)
}

// @Summary List work orders
// @Description Retrieve a paginated list of work orders with optional filters
// @Tags Work Orders
// @Produce json
// @Param page query int false "Page number"
// @Param per_page query int false "Items per page"
// @Param status query string false "Filter by status"
// @Param priority query string false "Filter by priority"
// @Param client_id query int false "Filter by client ID"
// @Param technician_id query int false "Filter by technician ID"
// @Param date_from query string false "Filter by start date"
// @Param date_to query string false "Filter by end date"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/work-orders [get]
func (h *WorkOrderHandler) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}
	filters := map[string]interface{}{
		"status":        c.Query("status"),
		"priority":      c.Query("priority"),
		"client_id":     c.Query("client_id"),
		"technician_id": c.Query("technician_id"),
		"date_from":     c.Query("date_from"),
		"date_to":       c.Query("date_to"),
	}
	orders, total, err := h.service.FindAll(page, perPage, filters)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.SuccessPaginated(c, "Work orders retrieved", orders, page, perPage, total)
}

// @Summary Update work order
// @Description Update an existing work order
// @Tags Work Orders
// @Accept json
// @Produce json
// @Param id path int true "Work order ID"
// @Param request body dto.UpdateWorkOrderRequest true "Updated work order data"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /api/v1/work-orders/{id} [put]
func (h *WorkOrderHandler) Update(c *gin.Context) {
	userID := getCurrentUserID(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	var req dto.UpdateWorkOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.service.Update(id, req, userID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusOK, "Work order updated", resp)
}

// @Summary Delete work order
// @Description Delete a work order by ID
// @Tags Work Orders
// @Produce json
// @Param id path int true "Work order ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/work-orders/{id} [delete]
func (h *WorkOrderHandler) Delete(c *gin.Context) {
	userID := getCurrentUserID(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	if err := h.service.Delete(id, userID); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusOK, "Work order deleted", nil)
}

// @Summary Assign technician
// @Description Assign a technician to a work order
// @Tags Work Orders
// @Accept json
// @Produce json
// @Param id path int true "Work order ID"
// @Param request body dto.AssignTechnicianRequest true "Technician ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /api/v1/work-orders/{id}/assign [post]
func (h *WorkOrderHandler) AssignTechnician(c *gin.Context) {
	userID := getCurrentUserID(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	var req dto.AssignTechnicianRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	if err := h.service.AssignTechnician(id, req.TechnicianID, userID); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusOK, "Technician assigned", nil)
}

// @Summary Change work order status
// @Description Change the status of a work order
// @Tags Work Orders
// @Accept json
// @Produce json
// @Param id path int true "Work order ID"
// @Param request body dto.ChangeStatusRequest true "New status"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /api/v1/work-orders/{id}/status [post]
func (h *WorkOrderHandler) ChangeStatus(c *gin.Context) {
	userID := getCurrentUserID(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	var req dto.ChangeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	if err := h.service.ChangeStatus(id, workorder.Status(req.Status), userID); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusOK, "Status changed", nil)
}

// @Summary Add note to work order
// @Description Add a note to a work order
// @Tags Work Orders
// @Accept json
// @Produce json
// @Param id path int true "Work order ID"
// @Param request body dto.AddNoteRequest true "Note text"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/work-orders/{id}/notes [post]
func (h *WorkOrderHandler) AddNote(c *gin.Context) {
	userID := getCurrentUserID(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	var req dto.AddNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	if err := h.service.AddNote(id, userID, req.Text); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusCreated, "Note added", nil)
}

// @Summary Add material to work order
// @Description Add a material with quantity to a work order
// @Tags Work Orders
// @Accept json
// @Produce json
// @Param id path int true "Work order ID"
// @Param request body dto.AddMaterialRequest true "Material ID and quantity"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/work-orders/{id}/materials [post]
func (h *WorkOrderHandler) AddMaterial(c *gin.Context) {
	userID := getCurrentUserID(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	var req dto.AddMaterialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	if err := h.service.AddMaterial(id, req.MaterialID, userID, req.Quantity); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusCreated, "Material added", nil)
}
