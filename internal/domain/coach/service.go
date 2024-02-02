package coach

import (
	"context"
	"errors"
)

// Service handles business logic related to coach entities.
type Service struct {
	coachRepository *Repository
}

// NewService creates a new Service instance.
func NewService(coachRepository *Repository) *Service {
	return &Service{
		coachRepository: coachRepository,
	}
}

// CreateCoach creates a new coach in the system.
func (s *Service) CreateCoach(ctx context.Context, name, specialty string) (*Coach, error) {
	if name == "" || specialty == "" {
		return nil, errors.New("name and specialty are required fields")
	}

	newCoach := NewCoach(name, specialty)

	if err := s.coachRepository.CreateCoach(ctx, newCoach); err != nil {
		return nil, err
	}

	return newCoach, nil
}

// GetCoachByID retrieves a coach by ID.
func (s *Service) GetCoachByID(ctx context.Context, coachID string) (*Coach, error) {
	return s.coachRepository.GetCoachByID(ctx, coachID)
}

// UpdateCoach updates an existing coach.
func (s *Service) UpdateCoach(ctx context.Context, coachID, name, specialty string) (*Coach, error) {
	if name == "" || specialty == "" {
		return nil, errors.New("name and specialty are required fields")
	}

	existingCoach, err := s.coachRepository.GetCoachByID(ctx, coachID)
	if err != nil {
		return nil, err
	}
	if existingCoach == nil {
		return nil, errors.New("coach not found")
	}

	existingCoach.Name = name
	existingCoach.Specialty = specialty

	if err := s.coachRepository.UpdateCoach(ctx, existingCoach); err != nil {
		return nil, err
	}

	return existingCoach, nil
}

// DeleteCoach deletes a coach by ID.
func (s *Service) DeleteCoach(ctx context.Context, coachID string) error {
	existingCoach, err := s.coachRepository.GetCoachByID(ctx, coachID)
	if err != nil {
		return err
	}
	if existingCoach == nil {
		return errors.New("coach not found")
	}

	return s.coachRepository.DeleteCoach(ctx, coachID)
}
