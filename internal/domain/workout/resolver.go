package workout

import (
	"context"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
	"time"
)

// Resolver handles GraphQL queries and mutations for the Log aggregate.
type Resolver struct {
	workoutService *Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(workoutService *Service) *Resolver {
	return &Resolver{
		workoutService: workoutService,
	}
}

// CreateWorkout is a GraphQL mutation to create a new workout entry.
func (r *Resolver) CreateWorkout(ctx context.Context, email, exercise, date string, sets, reps int, weight float64) (*graphql.Workout, error) {
	log.C(ctx).Infof("Creating Workout for user with email %q...", email)

	createdAt, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, err
	}

	workout, err := r.workoutService.CreateWorkout(ctx, email, exercise, sets, reps, weight, createdAt)
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
	workouts, err := r.workoutService.GetAllWorkouts(ctx, email, createdAt)
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

// GetLogQuery is a GraphQL query to retrieve a workout log entry by ID.
func (r *Resolver) GetLogQuery(ctx context.Context, workoutLogID string) (*Workout, error) {
	return r.workoutService.GetLogByID(ctx, workoutLogID)
}

// UpdateLogMutation is a GraphQL mutation to update an existing workout log entry.
func (r *Resolver) UpdateLogMutation(ctx context.Context, input UpdateLogInput) (*Workout, error) {
	return r.workoutService.UpdateLog(ctx, input.LogID, input.Exercise, input.Sets, input.Reps, input.Weight, input.LoggedAt)
}

// DeleteLogMutation is a GraphQL mutation to delete a workout log entry by ID.
func (r *Resolver) DeleteLogMutation(ctx context.Context, workoutLogID string) (string, error) {
	err := r.workoutService.DeleteLog(ctx, workoutLogID)
	if err != nil {
		return "", err
	}
	return "Workout log entry deleted successfully", nil
}

// CreateLogInput is the input structure for creating a new workout log entry.
type CreateLogInput struct {
	UserID   string    `json:"userId"`
	Exercise string    `json:"exercise"`
	Sets     int       `json:"sets"`
	Reps     int       `json:"reps"`
	Weight   float64   `json:"weight"`
	LoggedAt time.Time `json:"loggedAt"`
}

// UpdateLogInput is the input structure for updating an existing workout log entry.
type UpdateLogInput struct {
	LogID    string    `json:"workoutLogId"`
	Exercise string    `json:"exercise"`
	Sets     int       `json:"sets"`
	Reps     int       `json:"reps"`
	Weight   float64   `json:"weight"`
	LoggedAt time.Time `json:"loggedAt"`
}
