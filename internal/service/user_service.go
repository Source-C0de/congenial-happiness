package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/source-c0de/contacthub/internal/models"
	"github.com/source-c0de/contacthub/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetDashboardStats(ctx context.Context) (map[string]interface{}, error)
	ListUsers(ctx context.Context) ([]models.User, error)
	InviteUser(ctx context.Context, req *models.InviteUserRequest) (*models.User, error)
	UpdateRole(ctx context.Context, id uuid.UUID, role string) error
	SetUserStatus(ctx context.Context, id uuid.UUID, isActive bool) error
	SetPassword(ctx context.Context, userID uuid.UUID, req *models.SetPasswordRequest) error
	GetSessionSettings(ctx context.Context) (*models.SessionSettings, error)
	UpdateSessionSettings(ctx context.Context, timeoutMins int, logoutOnClose bool, updatedBy uuid.UUID) (*models.SessionSettings, error)
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

// InviteUser creates a real user account with a hashed password immediately.
func (s *userService) InviteUser(ctx context.Context, req *models.InviteUserRequest) (*models.User, error) {
	// Check if user already exists
	existing, _ := s.userRepo.GetByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	return s.userRepo.CreateUser(ctx, nil, req.Email, string(hashedPassword), req.Role)
}

func (s *userService) UpdateRole(ctx context.Context, id uuid.UUID, role string) error {
	return s.userRepo.UpdateRole(ctx, id, role)
}

func (s *userService) SetUserStatus(ctx context.Context, id uuid.UUID, isActive bool) error {
	return s.userRepo.SetUserStatus(ctx, id, isActive)
}

// SetPassword allows an admin to forcefully set any user's password.
func (s *userService) SetPassword(ctx context.Context, userID uuid.UUID, req *models.SetPasswordRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	return s.userRepo.SetUserPassword(ctx, userID, string(hashedPassword))
}

func (s *userService) GetSessionSettings(ctx context.Context) (*models.SessionSettings, error) {
	return s.userRepo.GetSessionSettings(ctx)
}

func (s *userService) UpdateSessionSettings(ctx context.Context, timeoutMins int, logoutOnClose bool, updatedBy uuid.UUID) (*models.SessionSettings, error) {
	return s.userRepo.UpdateSessionSettings(ctx, timeoutMins, logoutOnClose, updatedBy)
}
