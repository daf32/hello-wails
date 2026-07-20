package service

import (
	"context"
	"fmt"

	"hello-wails/internal/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (domain.User, error)
	List(ctx context.Context, limit int) ([]domain.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, id int) (domain.User, error) {
	if id <= 0 {
		return domain.User{}, domain.ErrInvalidID
	}

	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user %d: %w", id, err)
	}

	return user, nil
}

func (s *UserService) ListUsers(ctx context.Context, limit int) ([]domain.User, error) {
	if limit <= 0 {
		return nil, fmt.Errorf("%w: limit must be positive", domain.ErrInvalidID)
	}
	if limit > 100 {
		limit = 100
	}

	users, err := s.repo.List(ctx, limit)
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}

	return users, nil
}
