package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/source-c0de/contacthub/internal/models"
	"github.com/source-c0de/contacthub/internal/repository"
)

type EmployeeService interface {
	ListEmployees(ctx context.Context, filter *models.EmployeeFilter) (*models.EmployeeResponse, error)
	GetEmployee(ctx context.Context, id uuid.UUID) (*models.Employee, error)
	CreateEmployee(ctx context.Context, req *models.CreateEmployeeRequest) (*models.Employee, error)
	UpdateEmployee(ctx context.Context, id uuid.UUID, req *models.UpdateEmployeeRequest) (*models.Employee, error)
	DeleteEmployee(ctx context.Context, id uuid.UUID) error
}

type employeeService struct {
	empRepo *repository.EmployeeRepository
}

func NewEmployeeService(empRepo *repository.EmployeeRepository) EmployeeService {
	return &employeeService{empRepo: empRepo}
}

func (s *employeeService) ListEmployees(ctx context.Context, filter *models.EmployeeFilter) (*models.EmployeeResponse, error) {
	return s.empRepo.List(ctx, filter)
}

func (s *employeeService) GetEmployee(ctx context.Context, id uuid.UUID) (*models.Employee, error) {
	return s.empRepo.GetByID(ctx, id)
}

func (s *employeeService) CreateEmployee(ctx context.Context, req *models.CreateEmployeeRequest) (*models.Employee, error) {
	deptID, _ := uuid.Parse(req.DepartmentID)
	emp := &models.Employee{
		ID:           uuid.New(),
		FullName:     req.FullName,
		JobTitle:     req.JobTitle,
		DepartmentID: deptID,
		WorkEmail:    req.WorkEmail,
		Extension:    req.Extension,
	}
	if req.Mobile != nil {
		emp.Mobile = *req.Mobile
	}
	if req.OfficeLocation != nil {
		emp.OfficeLocation = *req.OfficeLocation
	}

	err := s.empRepo.Create(ctx, emp)
	return emp, err
}

func (s *employeeService) UpdateEmployee(ctx context.Context, id uuid.UUID, req *models.UpdateEmployeeRequest) (*models.Employee, error) {
	return s.empRepo.Update(ctx, id, req)
}

func (s *employeeService) DeleteEmployee(ctx context.Context, id uuid.UUID) error {
	return s.empRepo.Delete(ctx, id)
}
