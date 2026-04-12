package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/source-c0de/contacthub/internal/models"
	"github.com/source-c0de/contacthub/internal/repository"
)

type UserService interface {
	GetDashboardStats(ctx context.Context) (map[string]interface{}, error)
	ListUsers(ctx context.Context) ([]models.User, error)
	InviteUser(ctx context.Context, req *models.InviteUserRequest) error
	UpdateRole(ctx context.Context, id uuid.UUID, role string) error
	SetUserStatus(ctx context.Context, id uuid.UUID, isActive bool) error
}

type userService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetDashboardStats(ctx context.Context) (map[string]interface{}, error) {
	return s.userRepo.GetDashboardStats(ctx)
}

func (s *userService) ListUsers(ctx context.Context) ([]models.User, error) {
	return s.userRepo.ListUsers(ctx)
}

func (s *userService) InviteUser(ctx context.Context, req *models.InviteUserRequest) error {
	// Implementation would typically involve sending an email and creating a placeholder user
	return nil 
}

func (s *userService) UpdateRole(ctx context.Context, id uuid.UUID, role string) error {
	return s.userRepo.UpdateRole(ctx, id, role)
}

func (s *userService) SetUserStatus(ctx context.Context, id uuid.UUID, isActive bool) error {
	return s.userRepo.SetUserStatus(ctx, id, isActive)
}
