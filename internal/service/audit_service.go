package service

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	"github.com/source-c0de/contacthub/internal/models"
	"github.com/source-c0de/contacthub/internal/repository"
)

type AuditService interface {
	LogAction(ctx context.Context, logEntry *models.AuditLog) error
	GetLogs(ctx context.Context) ([]models.AuditLog, error)
	ExportLogs(ctx context.Context) ([]byte, error)
}

type auditService struct {
	auditRepo *repository.AuditRepository
}

func NewAuditService(auditRepo *repository.AuditRepository) AuditService {
	return &auditService{auditRepo: auditRepo}
}

func (s *auditService) LogAction(ctx context.Context, logEntry *models.AuditLog) error {
	return s.auditRepo.LogAction(ctx, logEntry)
}

func (s *auditService) GetLogs(ctx context.Context) ([]models.AuditLog, error) {
	return s.auditRepo.GetLogs(ctx)
}

func (s *auditService) ExportLogs(ctx context.Context) ([]byte, error) {
	logs, err := s.auditRepo.ExportLogs(ctx)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write header
	_ = writer.Write([]string{"ID", "UserID", "Action", "Entity", "EntityID", "OldData", "NewData", "IPAddress", "CreatedAt"})

	// Write data
	for _, log := range logs {
		_ = writer.Write([]string{
			fmt.Sprintf("%d", log.ID),
			log.UserID.String(),
			log.Action,
			log.Entity,
			log.EntityID,
			log.OldData,
			log.NewData,
			log.IPAddress,
			log.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	writer.Flush()

	return buf.Bytes(), nil
}
