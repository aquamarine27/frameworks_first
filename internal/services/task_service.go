package services

import (
	"frameworks_first/internal/domain"
	"frameworks_first/internal/errors"
)

type TaskService struct {
	repo *InMemoryRepository
}

func NewTaskService(repo *InMemoryRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAll() ([]*domain.TaskItem, error) {
	return s.repo.GetAll()
}

func (s *TaskService) GetByID(id int) (*domain.TaskItem, error) {
	item, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.ErrNotFound
	}
	return item, nil
}

func (s *TaskService) Create(item *domain.TaskItem) (*domain.TaskItem, error) {

	// validation
	if item.Name == "" {
		return nil, errors.ErrValidation.WithMessage("Name cannot be empty")
	}
	if item.Difficulty < 1 || item.Difficulty > 5 {
		return nil, errors.ErrValidation.WithMessage("Difficulty must be between 1 and 5")
	}
	if len(item.Description) > 500 {
		return nil, errors.ErrValidation.WithMessage("Description must be <= 500 characters")
	}

	return s.repo.Add(item)
}
