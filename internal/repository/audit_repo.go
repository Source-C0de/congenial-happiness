package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/source-c0de/contacthub/internal/models"
)

type AuditRepository struct {
	DB *sqlx.DB
}

func NewAuditRepository(db *sqlx.DB) *AuditRepository {
	return &AuditRepository{DB: db}
}

func (r *AuditRepository) LogAction(ctx context.Context, logEntry *models.AuditLog) error {
	query := `INSERT INTO audit_logs (user_id, action, entity, entity_id, old_data, new_data, ip_address, created_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())`
	_, err := r.DB.ExecContext(ctx, query, logEntry.UserID, logEntry.Action, logEntry.Entity, logEntry.EntityID, logEntry.OldData, logEntry.NewData, logEntry.IPAddress)
	return err
}

func (r *AuditRepository) GetLogs(ctx context.Context) ([]models.AuditLog, error) {
	var logs []models.AuditLog
	query := `SELECT id, user_id, action, entity, entity_id, old_data, new_data, ip_address, created_at FROM audit_logs ORDER BY created_at DESC LIMIT 500`
	err := r.DB.SelectContext(ctx, &logs, query)
	return logs, err
}

func (r *AuditRepository) ExportLogs(ctx context.Context) ([]models.AuditLog, error) {
	// For export, we might fetch more logs or format them differently, here just reusing GetLogs or a specific query
	var logs []models.AuditLog
	query := `SELECT id, user_id, action, entity, entity_id, old_data, new_data, ip_address, created_at FROM audit_logs ORDER BY created_at DESC`
	err := r.DB.SelectContext(ctx, &logs, query)
	return logs, err
}
