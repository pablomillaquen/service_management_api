package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/services"
	"github.com/pablomillaquen/speckit_golang_api/pkg/response"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary Create user
// @Description Create a new user (admin only)
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "User data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.service.Create(req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusCreated, "User created", resp)
}

// @Summary Get user by ID
// @Description Retrieve a single user by their ID (admin only)
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) FindByID(c *gin.Context) {
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
	response.Success(c, http.StatusOK, "User found", resp)
}

// @Summary List users
// @Description Retrieve a paginated list of users (admin only)
// @Tags Users
// @Produce json
// @Param page query int false "Page number"
// @Param per_page query int false "Items per page"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/users [get]
func (h *UserHandler) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}
	users, total, err := h.service.FindAll(page, perPage)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.SuccessPaginated(c, "Users retrieved", users, page, perPage, total)
}

// @Summary Update user
// @Description Update an existing user's data (admin only)
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserRequest true "Updated user data"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 409 {object} response.APIResponse
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid request", []string{err.Error()})
		return
	}
	resp, err := h.service.Update(id, req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusOK, "User updated", resp)
}

// @Summary Delete user
// @Description Delete a user by ID (admin only)
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "Invalid ID", nil)
		return
	}
	if err := h.service.Delete(id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(c, http.StatusOK, "User deleted", nil)
}
