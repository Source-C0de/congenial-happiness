package models

import (
	"time"
	"github.com/google/uuid"
)

type AuditLog struct {
	ID        int64     `db:"id"         json:"id"`
	UserID    uuid.UUID `db:"user_id"    json:"user_id"`
	Action    string    `db:"action"     json:"action"`
	Entity    string    `db:"entity"     json:"entity"`
	EntityID  string    `db:"entity_id"  json:"entity_id"`
	OldData   string    `db:"old_data"   json:"old_data"`
	NewData   string    `db:"new_data"   json:"new_data"`
	IPAddress string    `db:"ip_address" json:"ip_address"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
