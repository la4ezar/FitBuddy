package coach

import (
	"context"
)

// Service handles business logic related to coach entities.
type Service struct {
	repository *Repository
}

// NewService creates a new Service instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// GetAllCoaches gets all coaches
func (s *Service) GetAllCoaches(ctx context.Context) ([]*Coach, error) {
	return s.repository.GetAllCoaches(ctx)
}
