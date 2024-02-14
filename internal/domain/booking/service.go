package booking

import "context"

// Service handles business logic related to bookings.
type Service struct {
	repository *Repository
}

// NewService creates a new Service instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
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
