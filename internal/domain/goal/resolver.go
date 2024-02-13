package goal

import (
	"context"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
	"time"
)

// Resolver handles GraphQL queries and mutations for the Goal aggregate.
type Resolver struct {
	service *Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(service *Service) *Resolver {
	return &Resolver{
		service: service,
	}
}

// GetGoals is a GraphQL query to get all goals for user with given email.
func (r *Resolver) GetGoals(ctx context.Context, userEmail string) ([]*graphql.Goal, error) {
	log.C(ctx).Infof("Getting all goals for user with email %q...", userEmail)

	goals, err := r.service.GetGoals(ctx, userEmail)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully got all goals for user with email %q", userEmail)

	gqlGoals := make([]*graphql.Goal, 0, len(goals))
	for _, g := range goals {
		gqlGoals = append(gqlGoals, &graphql.Goal{
			ID:          g.ID,
			Name:        g.Name,
			Description: g.Description,
			StartDate:   g.StartDate.Format("2006-01-02 15:04:05"),
			EndDate:     g.EndDate.Format("2006-01-02 15:04:05"),
		})
	}

	return gqlGoals, nil
}

// CreateGoal is a GraphQL mutation to create a new fitness or wellness goal.
func (r *Resolver) CreateGoal(ctx context.Context, userEmail, name, description string, startDate, endDate time.Time) (*graphql.Goal, error) {
	log.C(ctx).Infof("Creating goal with name %q, description %q, start date %q and end date %q for user with email %q...", name, description, startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05"), userEmail)

	goal, err := r.service.CreateGoal(ctx, userEmail, name, description, startDate, endDate)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully created goal with name %q, description %q, start date %q and end date %q for user with email %q", name, description, startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05"), userEmail)

	return &graphql.Goal{
		ID:          goal.ID,
		Name:        goal.Name,
		Description: goal.Description,
		StartDate:   goal.StartDate.Format("2006-01-02 15:04:05"),
		EndDate:     goal.EndDate.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetGoalQuery is a GraphQL query to retrieve a fitness or wellness goal by ID.
func (r *Resolver) GetGoalQuery(ctx context.Context, goalID string) (*Goal, error) {
	return r.service.GetGoalByID(ctx, goalID)
}

// UpdateGoalMutation is a GraphQL mutation to update an existing fitness or wellness goal.
func (r *Resolver) UpdateGoalMutation(ctx context.Context, input UpdateGoalInput) (*Goal, error) {
	return r.service.UpdateGoal(ctx, input.GoalID, input.Title, input.Description, input.StartDate, input.EndDate)
}

// DeleteGoalMutation is a GraphQL mutation to delete a fitness or wellness goal by ID.
func (r *Resolver) DeleteGoalMutation(ctx context.Context, goalID string) (string, error) {
	err := r.service.DeleteGoal(ctx, goalID)
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
