package memory

import "sync"

type Repository struct {
	sync.RWMutex
	data map[string]string
}

func NewRepo() *Repository {
	return &Repository{data: make(map[string]string)}
}

// TODO: definir métodos para interactuar con DB
