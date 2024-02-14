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

// BookCoach books a coach
func (s *Service) BookCoach(ctx context.Context, email, coachName string) (bool, error) {
	return s.repository.BookCoach(ctx, email, coachName)
}

// UnbookCoach unbooks a coach
func (s *Service) UnbookCoach(ctx context.Context, email, coachName string) (bool, error) {
	return s.repository.UnbookCoach(ctx, email, coachName)
}

// IsCoachBookedByUser checks if a coach is booked by user
func (s *Service) IsCoachBookedByUser(ctx context.Context, coachName, email string) (bool, error) {
	return s.repository.IsCoachBookedByUser(ctx, coachName, email)
}

// IsCoachBooked checks if a coach is booked
func (s *Service) IsCoachBooked(ctx context.Context, coachName string) (bool, error) {
	return s.repository.IsCoachBooked(ctx, coachName)
}
