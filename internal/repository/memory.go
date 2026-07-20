package repository

import (
	"context"
	"fmt"
	"sync"

	"hello-wails/internal/domain"
)

type MemoryUserRepo struct {
	mu    sync.RWMutex
	users map[int]domain.User
}

func NewMemoryUserRepo() *MemoryUserRepo {
	users := make(map[int]domain.User, 20)
	for i := 1; i <= 20; i++ {
		users[i] = domain.User{
			ID:    i,
			Name:  fmt.Sprintf("user-%d", i),
			Email: fmt.Sprintf("user-%d@example.com", i),
		}
	}
	return &MemoryUserRepo{users: users}
}

func (r *MemoryUserRepo) GetByID(ctx context.Context, id int) (domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return domain.User{}, domain.ErrUserNotFound
	}
	return user, nil
}

func (r *MemoryUserRepo) List(ctx context.Context, limit int) ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]domain.User, 0, limit)
	for id := 1; id <= len(r.users) && len(result) < limit; id++ {
		if u, ok := r.users[id]; ok {
			result = append(result, u)
		}
	}
	return result, nil
}
