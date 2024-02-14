package exercise

import (
	"context"
)

// Service handles business logic related to exercise operations.
type Service struct {
	repository *Repository
}

// NewService creates a new Service instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// GetAllExercises retrieves all exercises.
func (s *Service) GetAllExercises(ctx context.Context) ([]*Exercise, error) {
	return s.repository.GetAllExercises(ctx)
}
