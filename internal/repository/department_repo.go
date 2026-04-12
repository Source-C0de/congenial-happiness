package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/source-c0de/contacthub/internal/models"
)

type DepartmentRepository struct {
	DB *sqlx.DB
}

func NewDepartmentRepository(db *sqlx.DB) *DepartmentRepository {
	return &DepartmentRepository{DB: db}
}

func (r *DepartmentRepository) List(ctx context.Context) ([]models.Department, error) {
	var depts []models.Department
	query := `SELECT id, name, code, created_at, updated_at FROM departments ORDER BY name ASC`
	err := r.DB.SelectContext(ctx, &depts, query)
	return depts, err
}

func (r *DepartmentRepository) Create(ctx context.Context, dept *models.Department) error {
	query := `INSERT INTO departments (id, name, code, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())`
	_, err := r.DB.ExecContext(ctx, query, dept.ID, dept.Name, dept.Code)
	return err
}

func (r *DepartmentRepository) Update(ctx context.Context, id uuid.UUID, req *models.UpdateDepartmentRequest) (*models.Department, error) {
	query := `UPDATE departments SET name = COALESCE($1, name), code = COALESCE($2, code), updated_at = NOW() WHERE id = $3 RETURNING id, name, code, created_at, updated_at`
	var dept models.Department
	err := r.DB.GetContext(ctx, &dept, query, req.Name, req.Code, id)
	return &dept, err
}

func (r *DepartmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM departments WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}
