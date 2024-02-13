package coach

import (
	"context"
	"errors"
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

func (s *Service) GetAllCoaches(ctx context.Context) ([]*Coach, error) {
	return s.repository.GetAllCoaches(ctx)
}

func (s *Service) BookCoach(ctx context.Context, email, coachName string) (bool, error) {
	return s.repository.BookCoach(ctx, email, coachName)
}

func (s *Service) UnbookCoach(ctx context.Context, email, coachName string) (bool, error) {
	return s.repository.UnbookCoach(ctx, email, coachName)
}

func (s *Service) IsCoachBookedByUser(ctx context.Context, coachName, email string) (bool, error) {
	return s.repository.IsCoachBookedByUser(ctx, coachName, email)
}

func (s *Service) IsCoachBooked(ctx context.Context, coachName string) (bool, error) {
	return s.repository.IsCoachBooked(ctx, coachName)
}

// CreateCoach creates a new coach in the system.
func (s *Service) CreateCoach(ctx context.Context, name, specialty string) (*Coach, error) {
	if name == "" || specialty == "" {
		return nil, errors.New("name and specialty are required fields")
	}

	newCoach := NewCoach(name, specialty, "")

	if err := s.repository.CreateCoach(ctx, newCoach); err != nil {
		return nil, err
	}

	return newCoach, nil
}

// GetCoachByID retrieves a coach by ID.
func (s *Service) GetCoachByID(ctx context.Context, coachID string) (*Coach, error) {
	return s.repository.GetCoachByID(ctx, coachID)
}

// UpdateCoach updates an existing coach.
func (s *Service) UpdateCoach(ctx context.Context, coachID, name, specialty string) (*Coach, error) {
	if name == "" || specialty == "" {
		return nil, errors.New("name and specialty are required fields")
	}

	existingCoach, err := s.repository.GetCoachByID(ctx, coachID)
	if err != nil {
		return nil, err
	}
	if existingCoach == nil {
		return nil, errors.New("coach not found")
	}

	existingCoach.Name = name
	existingCoach.Specialty = specialty

	if err := s.repository.UpdateCoach(ctx, existingCoach); err != nil {
		return nil, err
	}

	return existingCoach, nil
}

// DeleteCoach deletes a coach by ID.
func (s *Service) DeleteCoach(ctx context.Context, coachID string) error {
	existingCoach, err := s.repository.GetCoachByID(ctx, coachID)
	if err != nil {
		return err
	}
	if existingCoach == nil {
		return errors.New("coach not found")
	}

	return s.repository.DeleteCoach(ctx, coachID)
}
