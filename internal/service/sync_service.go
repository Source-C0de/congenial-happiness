package service

import (
	"context"

	"github.com/source-c0de/contacthub/internal/models"
	"github.com/source-c0de/contacthub/internal/repository"
)

type SyncService interface {
	GetSettings(ctx context.Context) ([]models.SyncSettings, error)
	ToggleSync(ctx context.Context, syncType string, isEnabled bool) error
	TriggerSync(ctx context.Context, syncType string) error
	GetLogs(ctx context.Context) ([]models.SyncLog, error)
}

type syncService struct {
	syncRepo *repository.SyncRepository
}

func NewSyncService(syncRepo *repository.SyncRepository) SyncService {
	return &syncService{syncRepo: syncRepo}
}

func (s *syncService) GetSettings(ctx context.Context) ([]models.SyncSettings, error) {
	return s.syncRepo.GetSettings(ctx)
}

func (s *syncService) ToggleSync(ctx context.Context, syncType string, isEnabled bool) error {
	return s.syncRepo.UpdateSettings(ctx, syncType, isEnabled)
}

func (s *syncService) TriggerSync(ctx context.Context, syncType string) error {
	// Dummy implementation for processing a sync job. In reality this might place a job in a queue
	logEntry := &models.SyncLog{
		Type:    syncType,
		Status:  "success",
		Message: "Manual sync triggered and completed successfully",
	}
	return s.syncRepo.LogSync(ctx, logEntry)
}

func (s *syncService) GetLogs(ctx context.Context) ([]models.SyncLog, error) {
	return s.syncRepo.GetLogs(ctx)
}
