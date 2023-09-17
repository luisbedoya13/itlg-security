package memory

import (
	"context"
	"itlg_security/create-token/internal/repository"
	"itlg_security/create-token/pkg/model"
	"sync"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.DdbUser
}

func New() *Repository {
	return &Repository{data: make(map[string]*model.DdbUser)}
}

func (r *Repository) GetUser(_ context.Context, email string) (*model.DdbUser, error) {
	r.RLock()
	defer r.RUnlock()
	m, ok := r.data[email]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil
}
