package goal

import (
	"context"
	"errors"
	"time"
)

// Service handles business logic related to fitness and wellness goals.
type Service struct {
	repository *Repository
}

// NewService creates a new Service instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// CreateGoal creates a new fitness or wellness goal.
func (s *Service) CreateGoal(ctx context.Context, userID, title, description string, startDate, endDate time.Time) (*Goal, error) {
	if userID == "" || title == "" || description == "" {
		return nil, errors.New("user ID, title, and description are required")
	}

	newGoal := NewGoal(userID, title, description, startDate, endDate)

	if err := s.repository.CreateGoal(ctx, newGoal); err != nil {
		return nil, err
	}

	return newGoal, nil
}

// GetGoalByID retrieves a fitness or wellness goal by ID.
func (s *Service) GetGoalByID(ctx context.Context, goalID string) (*Goal, error) {
	return s.repository.GetGoalByID(ctx, goalID)
}

// UpdateGoal updates an existing fitness or wellness goal.
func (s *Service) UpdateGoal(ctx context.Context, goalID, title, description string, startDate, endDate time.Time) (*Goal, error) {
	if title == "" || description == "" {
		return nil, errors.New("title and description are required")
	}

	existingGoal, err := s.repository.GetGoalByID(ctx, goalID)
	if err != nil {
		return nil, err
	}
	if existingGoal == nil {
		return nil, errors.New("goal not found")
	}

	existingGoal.Title = title
	existingGoal.Description = description
	existingGoal.StartDate = startDate
	existingGoal.EndDate = endDate

	if err := s.repository.UpdateGoal(ctx, existingGoal); err != nil {
		return nil, err
	}

	return existingGoal, nil
}

// DeleteGoal deletes a fitness or wellness goal by ID.
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
