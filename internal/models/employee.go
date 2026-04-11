package models


import (
	"time"
	"github.com/google/uuid"
)

type Employee struct {
	ID 			uuid.UUID  `db:"id"  json:"id"`
	FullName 	string     `db:"full_name"  json:"full_name"`
	JobTitle 	string     `db:"job_title"  json:"job_title"`
	DepartmentID uuid.UUID `db:"department_id"  json:"department_id"`
	DepartmentName string `db:"department_name"  json:"department_name"`
	Extension string 	   `db:"extension"  json:"extension"`
	WorkEmail 		string     `db:"work_email"  json:"work_email"`
	Mobile 		string     `db:"mobile"  json:"mobile"`
	OfficeLocation string     `db:"office_location"  json:"office_location"`
	IsActive 	bool       `db:"is_active"  json:"is_active"`
	CreatedAt 	time.Time  `db:"created_at"  json:"created_at"`
	UpdatedAt 	time.Time  `db:"updated_at"  json:"updated_at"`	
}

type CreateEmployeeRequest struct {
    FullName       string  `json:"full_name"       validate:"required,min=2,max=150"`
    JobTitle       string  `json:"job_title"       validate:"required,min=2,max=150"`
    DepartmentID   string  `json:"department_id"   validate:"required,uuid"`
    WorkEmail      string  `json:"work_email"      validate:"required,email"`
    Extension      string  `json:"extension"       validate:"required,min=3,max=10"`
    Mobile         *string `json:"mobile"`
    OfficeLocation *string `json:"office_location"`
}

type UpdateEmployeeRequest struct {
    FullName       *string `json:"full_name"       validate:"omitempty,min=2,max=150"`
    JobTitle       *string `json:"job_title"       validate:"omitempty,min=2,max=150"`
    DepartmentID   *string `json:"department_id"   validate:"omitempty,uuid"`
    WorkEmail      *string `json:"work_email"      validate:"omitempty,email"`
    Extension      *string `json:"extension"       validate:"omitempty,min=3,max=10"`
    Mobile         *string `json:"mobile"`
    OfficeLocation *string `json:"office_location"`
    IsActive       *bool   `json:"is_active"`
}

type EmployeeFilter struct {
	Search string `json:"search"`
	DepartmentID string `json:"department_id"`
	IsActive *bool `json:"is_active"`
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	SortBy string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
}

type EmployeeResponse struct {
    Data []Employee `json:"data"`
    Total int `json:"total"`
    Page int `json:"page"`
    PageSize int `json:"page_size"`
    TotalPages int `json:"total_pages"`
}