package goal

import (
	"context"
	"time"
)

// Resolver handles GraphQL queries and mutations for the Goal aggregate.
type Resolver struct {
	goalService *Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(goalService *Service) *Resolver {
	return &Resolver{
		goalService: goalService,
	}
}

// CreateGoalMutation is a GraphQL mutation to create a new fitness or wellness goal.
func (r *Resolver) CreateGoalMutation(ctx context.Context, input CreateGoalInput) (*Goal, error) {
	return r.goalService.CreateGoal(ctx, input.UserID, input.Title, input.Description, input.StartDate, input.EndDate)
}

// GetGoalQuery is a GraphQL query to retrieve a fitness or wellness goal by ID.
func (r *Resolver) GetGoalQuery(ctx context.Context, goalID string) (*Goal, error) {
	return r.goalService.GetGoalByID(ctx, goalID)
}

// UpdateGoalMutation is a GraphQL mutation to update an existing fitness or wellness goal.
func (r *Resolver) UpdateGoalMutation(ctx context.Context, input UpdateGoalInput) (*Goal, error) {
	return r.goalService.UpdateGoal(ctx, input.GoalID, input.Title, input.Description, input.StartDate, input.EndDate)
}

// DeleteGoalMutation is a GraphQL mutation to delete a fitness or wellness goal by ID.
func (r *Resolver) DeleteGoalMutation(ctx context.Context, goalID string) (string, error) {
	err := r.goalService.DeleteGoal(ctx, goalID)
	if err != nil {
		return "", err
	}
	return "Goal deleted successfully", nil
}

// CreateGoalInput is the input structure for creating a new fitness or wellness goal.
type CreateGoalInput struct {
	UserID      string    `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}

// UpdateGoalInput is the input structure for updating an existing fitness or wellness goal.
type UpdateGoalInput struct {
	GoalID      string    `json:"goalId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}
