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
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type InviteUserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Role  string `json:"role" binding:"required,oneof=admin user"`
}

type UpdateRoleRequest struct {
	Role string `json:"role" binding:"required,oneof=admin user"`
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
