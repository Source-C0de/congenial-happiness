package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID  `db:"id"  json:"id"`
	EmployeeId   *uuid.UUID `db:"employee_id"  json:"employee_id"`
	Email        string     `db:"email"  json:"email"`
	PasswordHash string     `db:"password_hash"  json:"-"`
	Role         string     `db:"role"  json:"role"`
	IsActive     bool       `db:"is_active"  json:"is_active"`
	LastLogin    *time.Time `db:"last_login_at"  json:"last_login"`
	CreatedAt    time.Time  `db:"created_at"  json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"  json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

type JWTClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}
