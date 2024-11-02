package inmemory

import (
	"sync"

	"github.com/kulinsky/assignment/internal/domain/assignment"
)

type Repository struct {
	storage map[assignment.ID]*assignment.Assignment
	mu      sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{
		storage: make(map[assignment.ID]*assignment.Assignment),
	}
}

func (r *Repository) Add(assignment *assignment.Assignment) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.storage[assignment.ID] = assignment

	return nil
}

func (r *Repository) Find(id assignment.ID) (*assignment.Assignment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	a, ok := r.storage[id]
	if !ok {
		return nil, assignment.ErrAssignmentNotFound
	}

	return a, nil
}
