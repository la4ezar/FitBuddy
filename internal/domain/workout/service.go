// internal/domain/workout/workout_service.go

package workout

import (
	"context"
	"errors"
	"github.com/FitBuddy/pkg/log"
	"time"
)

// Service handles business logic related to workout entries.
type Service struct {
	workoutRepository *Repository
}

// NewService creates a new Service instance.
func NewService(workoutRepository *Repository) *Service {
	return &Service{
		workoutRepository: workoutRepository,
	}
}

// CreateWorkout creates a new workout entry.
func (s *Service) CreateWorkout(ctx context.Context, email, exercise string, sets, reps int, weight float64, createdAt time.Time) (*Workout, error) {
	if email == "" || exercise == "" || sets <= 0 || reps <= 0 || weight <= 0 || createdAt == (time.Time{}) {
		return nil, errors.New("email, exercise, createdAt time, positive sets, positive reps, and positive weight are required")
	}

	workout := NewWorkout(email, exercise, sets, reps, weight, createdAt)

	if err := s.workoutRepository.CreateWorkout(ctx, workout); err != nil {
		log.C(ctx).Infof("Creating Workout for user with email %q...", err)
		return nil, err
	}

	return workout, nil
}

// GetAllWorkouts retrieves all workouts.
func (s *Service) GetAllWorkouts(ctx context.Context, email string, date time.Time) ([]*Workout, error) {
	return s.workoutRepository.GetAllWorkouts(ctx, email, date)
}

// GetLogByID retrieves a workout log entry by ID.
func (s *Service) GetLogByID(ctx context.Context, workoutLogID string) (*Workout, error) {
	return s.workoutRepository.GetLogByID(ctx, workoutLogID)
}

// UpdateLog updates an existing workout log entry.
func (s *Service) UpdateLog(ctx context.Context, workoutLogID, exercise string, sets, reps int, weight float64, loggedAt time.Time) (*Workout, error) {
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

	existingLog.ExerciseName = exercise
	existingLog.Sets = sets
	existingLog.Reps = reps
	existingLog.Weight = weight
	existingLog.CreatedAt = loggedAt

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
