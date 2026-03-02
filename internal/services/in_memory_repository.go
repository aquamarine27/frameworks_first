package services

import (
	"sync"
	"sync/atomic"

	"frameworks_first/internal/domain"
)

type InMemoryRepository struct {
	items  map[int]*domain.TaskItem
	nextID atomic.Int32
	mu     sync.RWMutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		items: make(map[int]*domain.TaskItem),
	}
}

func (r *InMemoryRepository) GetAll() ([]*domain.TaskItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]*domain.TaskItem, 0, len(r.items))
	for _, item := range r.items {
		list = append(list, item)
	}
	return list, nil
}

func (r *InMemoryRepository) GetByID(id int) (*domain.TaskItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, exists := r.items[id]
	if !exists {
		return nil, nil
	}
	return item, nil
}

func (r *InMemoryRepository) Add(item *domain.TaskItem) (*domain.TaskItem, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := int(r.nextID.Add(1))
	item.ID = id
	r.items[id] = item
	return item, nil
}
