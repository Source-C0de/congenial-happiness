package models

import (
	"time"
	"github.com/google/uuid"
)

type Department struct {
	ID        uuid.UUID `db:"id"         json:"id"`
	Name      string    `db:"name"       json:"name"`
	Code      string    `db:"code"       json:"code"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CreateDepartmentRequest struct {
	Name string `json:"name" binding:"required,min=2,max=100"`
	Code string `json:"code" binding:"required,min=2,max=20"`
}

type UpdateDepartmentRequest struct {
	Name *string `json:"name" binding:"omitempty,min=2,max=100"`
	Code *string `json:"code" binding:"omitempty,min=2,max=20"`
}
