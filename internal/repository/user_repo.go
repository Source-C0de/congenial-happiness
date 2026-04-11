package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/source-c0de/contacthub/internal/models"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, employee_id, email, password_hash, role, is_active, last_login_at, created_at, updated_at
	          FROM users WHERE email = $1`
	err := r.DB.GetContext(ctx, &user, query, email)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	query := `SELECT id, employee_id, email, password_hash, role, is_active, last_login_at, created_at, updated_at
	          FROM users WHERE id = $1`
	err := r.DB.GetContext(ctx, &user, query, id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateLastLogin(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE users SET last_login_at = NOW() WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}
