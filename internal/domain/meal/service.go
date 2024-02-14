package meal

import "context"

// Service handles business logic related to meal entries.
type Service struct {
	repository *Repository
}

// NewService creates a new Meal instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// GetAllMeals retrieves all meals.
func (s *Service) GetAllMeals(ctx context.Context) ([]*Meal, error) {
	return s.repository.GetAllMeals(ctx)
}
