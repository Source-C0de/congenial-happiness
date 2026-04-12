package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/source-c0de/contacthub/internal/models"
	"github.com/source-c0de/contacthub/internal/service"
)

type SyncHandler struct {
	SyncSvc service.SyncService
}

func NewSyncHandler(svc service.SyncService) *SyncHandler {
	return &SyncHandler{SyncSvc: svc}
}

// GetSettings handles GET /api/v1/admin/sync/settings
func (h *SyncHandler) GetSettings(c *gin.Context) {
	settings, err := h.SyncSvc.GetSettings(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, settings)
}

// ToggleSync handles PATCH /api/v1/admin/sync/settings/:type
func (h *SyncHandler) ToggleSync(c *gin.Context) {
	syncType := c.Param("type")
	var req models.ToggleSyncRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.SyncSvc.ToggleSync(c, syncType, req.IsEnabled)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sync settings updated"})
}

// RunSync handles POST /api/v1/admin/sync/run/:type
func (h *SyncHandler) RunSync(c *gin.Context) {
	syncType := c.Param("type")
	err := h.SyncSvc.TriggerSync(c, syncType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sync triggered successfully"})
}

// GetLogs handles GET /api/v1/admin/sync/logs
func (h *SyncHandler) GetLogs(c *gin.Context) {
	logs, err := h.SyncSvc.GetLogs(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}
