package nutrition

import (
	"context"
	"errors"
	"time"
)

// Service handles business logic related to nutrition log entries.
type Service struct {
	repository *Repository
}

// NewService creates a new NutritionService instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// CreateLog creates a new nutrition log entry.
func (s *Service) CreateLog(ctx context.Context, userID, meal, description string, calories int, loggedAt time.Time) (*Log, error) {
	if userID == "" || meal == "" || description == "" || calories <= 0 {
		return nil, errors.New("user ID, meal, description, and positive calories are required")
	}

	newLog := NewLog(userID, meal, description, calories, loggedAt)

	if err := s.repository.CreateLog(ctx, newLog); err != nil {
		return nil, err
	}

	return newLog, nil
}

// GetLogByID retrieves a nutrition log entry by ID.
func (s *Service) GetLogByID(ctx context.Context, nutritionLogID string) (*Log, error) {
	return s.repository.GetLogByID(ctx, nutritionLogID)
}

// UpdateLog updates an existing nutrition log entry.
func (s *Service) UpdateLog(ctx context.Context, nutritionLogID, meal, description string, calories int, loggedAt time.Time) (*Log, error) {
	if meal == "" || description == "" || calories <= 0 {
		return nil, errors.New("meal, description, and positive calories are required")
	}

	existingLog, err := s.repository.GetLogByID(ctx, nutritionLogID)
	if err != nil {
		return nil, err
	}
	if existingLog == nil {
		return nil, errors.New("nutrition log entry not found")
	}

	existingLog.Meal = meal
	existingLog.Description = description
	existingLog.Calories = calories
	existingLog.LoggedAt = loggedAt

	if err := s.repository.UpdateLog(ctx, existingLog); err != nil {
		return nil, err
	}

	return existingLog, nil
}

// DeleteLog deletes a nutrition log entry by ID.
func (s *Service) DeleteLog(ctx context.Context, nutritionLogID string) error {
	existingLog, err := s.repository.GetLogByID(ctx, nutritionLogID)
	if err != nil {
		return err
	}
	if existingLog == nil {
		return errors.New("nutrition log entry not found")
	}

	return s.repository.DeleteLog(ctx, nutritionLogID)
}
