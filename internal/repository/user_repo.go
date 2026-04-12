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

func (r *UserRepository) UpdatePassword(ctx context.Context, id uuid.UUID, hashedPassword string) error {
	query := `UPDATE users SET password_hash = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, hashedPassword, id)
	return err
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	query := `SELECT id, employee_id, email, password_hash, role, is_active, last_login_at, created_at, updated_at FROM users`
	err := r.DB.SelectContext(ctx, &users, query)
	return users, err
}

func (r *UserRepository) UpdateRole(ctx context.Context, id uuid.UUID, role string) error {
	query := `UPDATE users SET role = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, role, id)
	return err
}

func (r *UserRepository) SetUserStatus(ctx context.Context, id uuid.UUID, isActive bool) error {
	query := `UPDATE users SET is_active = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, isActive, id)
	return err
}

func (r *UserRepository) GetDashboardStats(ctx context.Context) (map[string]interface{}, error) {
	var totalUsers int
	var activeEmployees int
	var totalDepartments int

	err := r.DB.GetContext(ctx, &totalUsers, "SELECT COUNT(*) FROM users")
	if err != nil {
		return nil, err
	}
	err = r.DB.GetContext(ctx, &activeEmployees, "SELECT COUNT(*) FROM employees WHERE is_active = true")
	if err != nil {
		return nil, err
	}
	err = r.DB.GetContext(ctx, &totalDepartments, "SELECT COUNT(*) FROM departments")
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total_users":       totalUsers,
		"active_employees":  activeEmployees,
		"total_departments": totalDepartments,
	}, nil
}

func (r *UserRepository) UpdateLastLogin(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE users SET last_login_at = NOW() WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}
