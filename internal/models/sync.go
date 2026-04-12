package models

import (
	"time"
)

type SyncSettings struct {
	Type      string    `json:"type"`
	IsEnabled bool      `json:"is_enabled"`
	LastRun   *time.Time `json:"last_run"`
}

type SyncLog struct {
	ID        int64     `json:"id"`
	Type      string    `json:"type"`
	Status    string    `json:"status"` // success, failure
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type ToggleSyncRequest struct {
	IsEnabled bool `json:"is_enabled"`
}
