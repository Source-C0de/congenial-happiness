package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/source-c0de/contacthub/internal/models"
)

type SyncRepository struct {
	DB *sqlx.DB
}

func NewSyncRepository(db *sqlx.DB) *SyncRepository {
	return &SyncRepository{DB: db}
}

func (r *SyncRepository) GetSettings(ctx context.Context) ([]models.SyncSettings, error) {
	// Dummy implementation for now, assuming settings are stored in db or fetched from config
	return []models.SyncSettings{
		{Type: "ldap", IsEnabled: true},
		{Type: "ad", IsEnabled: false},
	}, nil
}

func (r *SyncRepository) UpdateSettings(ctx context.Context, syncType string, isEnabled bool) error {
	// Implementation would update db table sync_settings
	query := `UPDATE sync_settings SET is_enabled = $1 WHERE type = $2`
	_, err := r.DB.ExecContext(ctx, query, isEnabled, syncType)
	return err
}

func (r *SyncRepository) LogSync(ctx context.Context, logEntry *models.SyncLog) error {
	query := `INSERT INTO sync_logs (type, status, message, created_at) VALUES ($1, $2, $3, NOW())`
	_, err := r.DB.ExecContext(ctx, query, logEntry.Type, logEntry.Status, logEntry.Message)
	return err
}

func (r *SyncRepository) GetLogs(ctx context.Context) ([]models.SyncLog, error) {
	var logs []models.SyncLog
	query := `SELECT id, type, status, message, created_at FROM sync_logs ORDER BY created_at DESC LIMIT 100`
	err := r.DB.SelectContext(ctx, &logs, query)
	return logs, err
}
