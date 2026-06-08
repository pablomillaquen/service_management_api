package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/services"
	"github.com/pablomillaquen/speckit_golang_api/pkg/response"
)

type ClientHandler struct {
	service *services.ClientService
}

func NewClientHandler(service *services.ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}

// @Summary Create client
// @Description Register a new client
// @Tags Clients
// @Accept json
// @Produce json
// @Param request body dto.CreateClientRequest true "Client data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/clients [post]
func (h *ClientHandler) Create(c *gin.Context) {
	var req dto.CreateClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.service.Create(req)
	if err != nil {
		response.Conflict(c, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Client created", resp)
}

// @Summary Get client by ID
// @Description Retrieve a single client by their ID
// @Tags Clients
// @Produce json
// @Param id path int true "Client ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/v1/clients/{id} [get]
func (h *ClientHandler) FindByID(c *gin.Context) {
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
	response.Success(c, http.StatusOK, "Client found", resp)
}

// @Summary List clients
// @Description Retrieve a paginated list of clients with optional search
// @Tags Clients
// @Produce json
// @Param page query int false "Page number"
// @Param per_page query int false "Items per page"
// @Param search query string false "Search term"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/clients [get]
func (h *ClientHandler) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}
	search := c.Query("search")
	clients, total, err := h.service.FindAll(page, perPage, search)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.SuccessPaginated(c, "Clients retrieved", clients, page, perPage, total)
}

// @Summary Update client
// @Description Update an existing client's data
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path int true "Client ID"
// @Param request body dto.UpdateClientRequest true "Updated client data"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/clients/{id} [put]
func (h *ClientHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	var req dto.UpdateClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.service.Update(id, req)
	if err != nil {
		response.Conflict(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Client updated", resp)
}

// @Summary Delete client
// @Description Delete a client by ID
// @Tags Clients
// @Produce json
// @Param id path int true "Client ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/clients/{id} [delete]
func (h *ClientHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	if err := h.service.Delete(id); err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Client deleted", nil)
}
