package repository

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/source-c0de/contacthub/internal/models"
)

type EmployeeRepository struct {
	DB *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) List(ctx context.Context, filter *models.EmployeeFilter) (*models.EmployeeResponse, error) {
	queryBuilder := []string{"SELECT e.*, d.name as department_name FROM employees e LEFT JOIN departments d ON e.department_id = d.id WHERE 1=1"}
	args := []interface{}{}
	argCounter := 1

	if filter.Search != "" {
		queryBuilder = append(queryBuilder, fmt.Sprintf("(e.full_name ILIKE $%d OR e.work_email ILIKE $%d OR e.mobile ILIKE $%d)", argCounter, argCounter, argCounter))
		args = append(args, "%"+filter.Search+"%")
		argCounter++
	}

	if filter.DepartmentID != "" {
		queryBuilder = append(queryBuilder, fmt.Sprintf("e.department_id = $%d", argCounter))
		args = append(args, filter.DepartmentID)
		argCounter++
	}

	if filter.IsActive != nil {
		queryBuilder = append(queryBuilder, fmt.Sprintf("e.is_active = $%d", argCounter))
		args = append(args, *filter.IsActive)
		argCounter++
	}

	query := strings.Join(queryBuilder, " AND ")

	// Count total
	countQuery := "SELECT COUNT(*) FROM (" + query + ") as c"
	var total int
	err := r.DB.GetContext(ctx, &total, countQuery, args...)
	if err != nil {
		return nil, err
	}

	// Pagination & Sorting
	if filter.SortBy != "" {
		sortOrder := "ASC"
		if filter.SortOrder == "desc" || filter.SortOrder == "DESC" {
			sortOrder = "DESC"
		}
		// Basic sanitization
		allowedSorts := map[string]bool{"full_name": true, "created_at": true, "department_name": true}
		if allowedSorts[filter.SortBy] {
			query += fmt.Sprintf(" ORDER BY %s %s", filter.SortBy, sortOrder)
		} else {
			query += " ORDER BY e.created_at DESC"
		}
	} else {
		query += " ORDER BY e.created_at DESC"
	}

	page := filter.Page
	if page < 1 {
		page = 1
	}
	pageSize := filter.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCounter, argCounter+1)
	args = append(args, pageSize, offset)

	var employees []models.Employee
	err = r.DB.SelectContext(ctx, &employees, query, args...)
	if err != nil {
		return nil, err
	}

	if employees == nil {
		employees = []models.Employee{}
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &models.EmployeeResponse{
		Data:       employees,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

func (r *EmployeeRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Employee, error) {
	var emp models.Employee
	query := `SELECT e.*, d.name as department_name FROM employees e LEFT JOIN departments d ON e.department_id = d.id WHERE e.id = $1`
	err := r.DB.GetContext(ctx, &emp, query, id)
	return &emp, err
}

func (r *EmployeeRepository) Create(ctx context.Context, emp *models.Employee) error {
	query := `INSERT INTO employees (id, full_name, job_title, department_id, extension, work_email, mobile, office_location, is_active, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, true, NOW(), NOW())`
	_, err := r.DB.ExecContext(ctx, query, emp.ID, emp.FullName, emp.JobTitle, emp.DepartmentID, emp.Extension, emp.WorkEmail, emp.Mobile, emp.OfficeLocation)
	return err
}

func (r *EmployeeRepository) Update(ctx context.Context, id uuid.UUID, req *models.UpdateEmployeeRequest) (*models.Employee, error) {
	// Simple update implementation, mapping values
	query := `UPDATE employees SET 
		full_name = COALESCE($1, full_name),
		job_title = COALESCE($2, job_title),
		department_id = COALESCE($3, department_id),
		work_email = COALESCE($4, work_email),
		extension = COALESCE($5, extension),
		mobile = COALESCE($6, mobile),
		office_location = COALESCE($7, office_location),
		is_active = COALESCE($8, is_active),
		updated_at = NOW()
		WHERE id = $9 RETURNING *`

	var emp models.Employee
	err := r.DB.GetContext(ctx, &emp, query, req.FullName, req.JobTitle, req.DepartmentID, req.WorkEmail, req.Extension, req.Mobile, req.OfficeLocation, req.IsActive, id)
	return &emp, err
}

func (r *EmployeeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM employees WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}
