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
	PhotoURL 	*string    `db:"photo_url"  json:"photo_url"`
	IsActive 	bool       `db:"is_active"  json:"is_active"`
	OnlineStatus string    `db:"online_status"  json:"online_status"`
	CreatedAt 	time.Time  `db:"created_at"  json:"created_at"`
	UpdatedAt 	time.Time  `db:"updated_at"  json:"updated_at"`
}

type CreateEmployeeRequest struct {
    FullName       string  `json:"full_name"       binding:"required,min=2,max=150"`
    JobTitle       string  `json:"job_title"       binding:"required,min=2,max=150"`
    DepartmentID   string  `json:"department_id"   binding:"required,uuid"`
    WorkEmail      string  `json:"work_email"      binding:"required,email"`
    Extension      string  `json:"extension"       binding:"required,min=3,max=10"`
    Mobile         *string `json:"mobile"`
    OfficeLocation *string `json:"office_location"`
}

type UpdateEmployeeRequest struct {
    FullName       *string `json:"full_name"       binding:"omitempty,min=2,max=150"`
    JobTitle       *string `json:"job_title"       binding:"omitempty,min=2,max=150"`
    DepartmentID   *string `json:"department_id"   binding:"omitempty,uuid"`
    WorkEmail      *string `json:"work_email"      binding:"omitempty,email"`
    Extension      *string `json:"extension"       binding:"omitempty,min=3,max=10"`
    Mobile         *string `json:"mobile"`
    OfficeLocation *string `json:"office_location"`
    IsActive       *bool   `json:"is_active"`
}

type EmployeeFilter struct {
	Search        string `form:"search"`
	DepartmentID  string `form:"department_id"`
	IsActive     *bool   `form:"is_active"`
	Page         int     `form:"page,default=1"`
	PageSize     int     `form:"page_size,default=10"`
	SortBy       string  `form:"sort_by"`
	SortOrder    string  `form:"sort_order"`
}

type EmployeeResponse struct {
    Data []Employee `json:"data"`
    Total int `json:"total"`
    Page int `json:"page"`
    PageSize int `json:"page_size"`
    TotalPages int `json:"total_pages"`
}