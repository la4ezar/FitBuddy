package workout

import (
	"context"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
	"time"
)

// Resolver handles GraphQL queries and mutations for the Log aggregate.
type Resolver struct {
	service *Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(service *Service) *Resolver {
	return &Resolver{
		service: service,
	}
}

// CreateWorkout is a GraphQL mutation to create a new workout entry.
func (r *Resolver) CreateWorkout(ctx context.Context, email, exercise, date string, sets, reps int, weight float64) (*graphql.Workout, error) {
	log.C(ctx).Infof("Creating Workout for user with email %q...", email)

	createdAt, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, err
	}

	workout, err := r.service.CreateWorkout(ctx, email, exercise, sets, reps, weight, createdAt)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully created workout for user with email %q", email)

	gqlWorkout := &graphql.Workout{
		ID:           workout.ID,
		UserEmail:    workout.UserEmail,
		ExerciseName: workout.ExerciseName,
		Sets:         workout.Sets,
		Reps:         workout.Reps,
		Weight:       workout.Weight,
		Date:         workout.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return gqlWorkout, nil
}

// GetAllWorkouts is a GraphQL query to retrieve all workouts.
func (r *Resolver) GetAllWorkouts(ctx context.Context, email, date string) ([]*graphql.Workout, error) {
	log.C(ctx).Info("Getting all workouts...")
	createdAt, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, err
	}
	workouts, err := r.service.GetAllWorkouts(ctx, email, createdAt)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Info("Successfully got all workouts")

	gqlWorkouts := make([]*graphql.Workout, 0, len(workouts))
	for _, w := range workouts {
		gqlWorkouts = append(gqlWorkouts, &graphql.Workout{
			ID:           w.ID,
			UserEmail:    w.UserEmail,
			ExerciseName: w.ExerciseName,
			Sets:         w.Sets,
			Reps:         w.Reps,
			Weight:       w.Weight,
			Date:         w.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return gqlWorkouts, nil
}
