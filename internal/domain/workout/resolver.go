package workout

import (
	"context"
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

// CreateLogMutation is a GraphQL mutation to create a new workout log entry.
func (r *Resolver) CreateLogMutation(ctx context.Context, input CreateLogInput) (*Log, error) {
	return r.workoutService.CreateLog(ctx, input.UserID, input.Exercise, input.Sets, input.Reps, input.Weight, input.LoggedAt)
}

// GetLogQuery is a GraphQL query to retrieve a workout log entry by ID.
func (r *Resolver) GetLogQuery(ctx context.Context, workoutLogID string) (*Log, error) {
	return r.workoutService.GetLogByID(ctx, workoutLogID)
}

// UpdateLogMutation is a GraphQL mutation to update an existing workout log entry.
func (r *Resolver) UpdateLogMutation(ctx context.Context, input UpdateLogInput) (*Log, error) {
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
