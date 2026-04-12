package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/source-c0de/contacthub/internal/models"
	"github.com/source-c0de/contacthub/internal/service"
)

type AdminHandler struct {
	AdminSvc service.UserService // Reusing UserService for admin operations
}

func NewAdminHandler(svc service.UserService) *AdminHandler {
	return &AdminHandler{AdminSvc: svc}
}

// GetDashboardStats handles GET /api/v1/admin/dashboard
func (h *AdminHandler) GetDashboardStats(c *gin.Context) {
	stats, err := h.AdminSvc.GetDashboardStats(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

// ListUsers handles GET /api/v1/admin/users
func (h *AdminHandler) ListUsers(c *gin.Context) {
	users, err := h.AdminSvc.ListUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// InviteUser handles POST /api/v1/admin/users/invite
func (h *AdminHandler) InviteUser(c *gin.Context) {
	var req models.InviteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.AdminSvc.InviteUser(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user invited successfully"})
}

// ChangeRole handles PATCH /api/v1/admin/users/:id/role
func (h *AdminHandler) ChangeRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req models.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.AdminSvc.UpdateRole(c, id, req.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role updated successfully"})
}

// DeactivateUser handles PATCH /api/v1/admin/users/:id/deactivate
func (h *AdminHandler) DeactivateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	err = h.AdminSvc.SetUserStatus(c, id, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deactivated successfully"})
}
