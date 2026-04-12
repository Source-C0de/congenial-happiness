package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/source-c0de/contacthub/internal/models"
	"github.com/source-c0de/contacthub/internal/repository"
)

type DepartmentService interface {
	ListDepartments(ctx context.Context) ([]models.Department, error)
	CreateDepartment(ctx context.Context, req *models.CreateDepartmentRequest) (*models.Department, error)
	UpdateDepartment(ctx context.Context, id uuid.UUID, req *models.UpdateDepartmentRequest) (*models.Department, error)
	DeleteDepartment(ctx context.Context, id uuid.UUID) error
}

type departmentService struct {
	deptRepo *repository.DepartmentRepository
}

func NewDepartmentService(deptRepo *repository.DepartmentRepository) DepartmentService {
	return &departmentService{deptRepo: deptRepo}
}

func (s *departmentService) ListDepartments(ctx context.Context) ([]models.Department, error) {
	return s.deptRepo.List(ctx)
}

func (s *departmentService) CreateDepartment(ctx context.Context, req *models.CreateDepartmentRequest) (*models.Department, error) {
	dept := &models.Department{
		ID:   uuid.New(),
		Name: req.Name,
		Code: req.Code,
	}
	err := s.deptRepo.Create(ctx, dept)
	return dept, err
}

func (s *departmentService) UpdateDepartment(ctx context.Context, id uuid.UUID, req *models.UpdateDepartmentRequest) (*models.Department, error) {
	return s.deptRepo.Update(ctx, id, req)
}

func (s *departmentService) DeleteDepartment(ctx context.Context, id uuid.UUID) error {
	return s.deptRepo.Delete(ctx, id)
}
