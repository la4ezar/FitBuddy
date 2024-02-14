package goal

import (
	"context"
	"errors"
	"time"
)

// Service handles business logic related to fitness goals.
type Service struct {
	repository *Repository
}

// NewService creates a new Service instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// CreateGoal creates a new fitness goal.
func (s *Service) CreateGoal(ctx context.Context, userEmail, name, description string, startDate, endDate time.Time) (*Goal, error) {
	if userEmail == "" {
		return nil, errors.New("user email should not be empty")
	}

	if name == "" || description == "" {
		return nil, errors.New("user ID, title, and description are required")
	}

	newGoal := NewGoal(name, description, startDate, endDate)

	if err := s.repository.CreateGoal(ctx, userEmail, newGoal); err != nil {
		return nil, err
	}

	return newGoal, nil
}

// GetGoals retrieves all goals for user with email
func (s *Service) GetGoals(ctx context.Context, userEmail string) ([]*Goal, error) {
	if userEmail == "" {
		return nil, errors.New("user email should not be empty")
	}

	return s.repository.GetGoalsByEmail(ctx, userEmail)
}

// DeleteGoal deletes a fitness goal by ID.
func (s *Service) DeleteGoal(ctx context.Context, goalID string) error {
	existingGoal, err := s.repository.GetGoalByID(ctx, goalID)
	if err != nil {
		return err
	}
	if existingGoal == nil {
		return errors.New("goal not found")
	}

	return s.repository.DeleteGoal(ctx, goalID)
}

// CompleteGoal deletes a fitness goal by ID.
func (s *Service) CompleteGoal(ctx context.Context, goalID string) error {
	existingGoal, err := s.repository.GetGoalByID(ctx, goalID)
	if err != nil {
		return err
	}
	if existingGoal == nil {
		return errors.New("goal not found")
	}
	if existingGoal.Completed {
		return nil
	}

	return s.repository.CompleteGoalByID(ctx, goalID)
}
