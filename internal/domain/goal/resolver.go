package goal

import (
	"context"
	"github.com/FitBuddy/internal/domain/leaderboard"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
	"time"
)

// Resolver handles GraphQL queries and mutations for the Goal aggregate.
type Resolver struct {
	service            *Service
	leaderboardService *leaderboard.Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(service *Service, leaderboardService *leaderboard.Service) *Resolver {
	return &Resolver{
		service:            service,
		leaderboardService: leaderboardService,
	}
}

// GetGoals is a gets all goals for user with given email.
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
			Completed:   g.Completed,
		})
	}

	return gqlGoals, nil
}

// CreateGoal is a creates a new fitness goal.
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

// DeleteGoal is a deletes a goal.
func (r *Resolver) DeleteGoal(ctx context.Context, goalID string) (bool, error) {
	log.C(ctx).Infof("Deleting goal with ID %q...", goalID)

	err := r.service.DeleteGoal(ctx, goalID)
	if err != nil {
		return false, err
	}
	log.C(ctx).Infof("Successfully deleted goal with ID %q...", goalID)

	return true, nil
}

// CompleteGoal is completes a goal.
func (r *Resolver) CompleteGoal(ctx context.Context, userEmail, goalID string) (bool, error) {
	log.C(ctx).Infof("Completing goal with ID %q...", goalID)

	err := r.service.CompleteGoal(ctx, goalID)
	if err != nil {
		return false, err
	}
	log.C(ctx).Infof("Successfully completed goal with ID %q...", goalID)

	log.C(ctx).Infof("Adding score to user with email %q...", userEmail)
	err = r.leaderboardService.AddScore(ctx, userEmail, 1)
	if err != nil {
		return false, err
	}
	log.C(ctx).Infof("Successfully added score to email %q...", userEmail)

	return true, nil
}
