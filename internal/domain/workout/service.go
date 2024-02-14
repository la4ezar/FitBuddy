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
	repository *Repository
}

// NewService creates a new Service instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// CreateWorkout creates a new workout entry.
func (s *Service) CreateWorkout(ctx context.Context, email, exercise string, sets, reps int, weight float64, createdAt time.Time) (*Workout, error) {
	if email == "" || exercise == "" || sets <= 0 || reps <= 0 || weight <= 0 || createdAt == (time.Time{}) {
		return nil, errors.New("email, exercise, createdAt time, positive sets, positive reps, and positive weight are required")
	}

	workout := NewWorkout(email, exercise, sets, reps, weight, createdAt)

	if err := s.repository.CreateWorkout(ctx, workout); err != nil {
		log.C(ctx).Infof("Creating Workout for user with email %q...", err)
		return nil, err
	}

	return workout, nil
}

// DeleteWorkout deletes a workout entry.
func (s *Service) DeleteWorkout(ctx context.Context, workoutID string) error {
	if workoutID == "" {
		return errors.New("workout ID is required")
	}

	if err := s.repository.DeleteWorkout(ctx, workoutID); err != nil {
		log.C(ctx).Infof("Deleting Workout with ID %q: %v", workoutID, err)
		return err
	}

	return nil
}

// GetAllWorkouts retrieves all workouts.
func (s *Service) GetAllWorkouts(ctx context.Context, email string, date time.Time) ([]*Workout, error) {
	return s.repository.GetAllWorkouts(ctx, email, date)
}
