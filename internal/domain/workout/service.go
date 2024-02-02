// internal/domain/workout/workout_service.go

package workout

import (
	"context"
	"errors"
	"time"
)

// Service handles business logic related to workout log entries.
type Service struct {
	workoutRepository *Repository
}

// NewService creates a new Service instance.
func NewService(workoutRepository *Repository) *Service {
	return &Service{
		workoutRepository: workoutRepository,
	}
}

// CreateLog creates a new workout log entry.
func (s *Service) CreateLog(ctx context.Context, userID, exercise string, sets, reps int, weight float64, loggedAt time.Time) (*Log, error) {
	if userID == "" || exercise == "" || sets <= 0 || reps <= 0 || weight <= 0 {
		return nil, errors.New("user ID, exercise, positive sets, positive reps, and positive weight are required")
	}

	newLog := NewLog(userID, exercise, sets, reps, weight, loggedAt)

	if err := s.workoutRepository.CreateLog(ctx, newLog); err != nil {
		return nil, err
	}

	return newLog, nil
}

// GetLogByID retrieves a workout log entry by ID.
func (s *Service) GetLogByID(ctx context.Context, workoutLogID string) (*Log, error) {
	return s.workoutRepository.GetLogByID(ctx, workoutLogID)
}

// UpdateLog updates an existing workout log entry.
func (s *Service) UpdateLog(ctx context.Context, workoutLogID, exercise string, sets, reps int, weight float64, loggedAt time.Time) (*Log, error) {
	if exercise == "" || sets <= 0 || reps <= 0 || weight <= 0 {
		return nil, errors.New("exercise, positive sets, positive reps, and positive weight are required")
	}

	existingLog, err := s.workoutRepository.GetLogByID(ctx, workoutLogID)
	if err != nil {
		return nil, err
	}
	if existingLog == nil {
		return nil, errors.New("workout log entry not found")
	}

	existingLog.Exercise = exercise
	existingLog.Sets = sets
	existingLog.Reps = reps
	existingLog.Weight = weight
	existingLog.LoggedAt = loggedAt

	if err := s.workoutRepository.UpdateLog(ctx, existingLog); err != nil {
		return nil, err
	}

	return existingLog, nil
}

// DeleteLog deletes a workout log entry by ID.
func (s *Service) DeleteLog(ctx context.Context, workoutLogID string) error {
	existingLog, err := s.workoutRepository.GetLogByID(ctx, workoutLogID)
	if err != nil {
		return err
	}
	if existingLog == nil {
		return errors.New("workout log entry not found")
	}

	return s.workoutRepository.DeleteLog(ctx, workoutLogID)
}
