package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/source-c0de/contacthub/internal/service"
)

type AuditHandler struct {
	AuditSvc service.AuditService
}

func NewAuditHandler(svc service.AuditService) *AuditHandler {
	return &AuditHandler{AuditSvc: svc}
}

// GetAuditLogs handles GET /api/v1/admin/audit-logs
func (h *AuditHandler) GetAuditLogs(c *gin.Context) {
	logs, err := h.AuditSvc.GetLogs(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

// ExportAuditLogs handles GET /api/v1/admin/audit-logs/export
func (h *AuditHandler) ExportAuditLogs(c *gin.Context) {
	data, err := h.AuditSvc.ExportLogs(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=audit-logs.csv")
	c.Data(http.StatusOK, "text/csv", data)
}
