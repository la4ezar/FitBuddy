package coach

import (
	"context"
)

// Resolver handles GraphQL queries and mutations for the Coach aggregate.
type Resolver struct {
	coachService *Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(coachService *Service) *Resolver {
	return &Resolver{
		coachService: coachService,
	}
}

// CreateCoachMutation is a GraphQL mutation to create a new coach.
func (r *Resolver) CreateCoachMutation(ctx context.Context, input CreateCoachInput) (*Coach, error) {
	return r.coachService.CreateCoach(ctx, input.Name, input.Specialty)
}

// GetCoachQuery is a GraphQL query to retrieve a coach by ID.
func (r *Resolver) GetCoachQuery(ctx context.Context, coachID string) (*Coach, error) {
	return r.coachService.GetCoachByID(ctx, coachID)
}

// UpdateCoachMutation is a GraphQL mutation to update an existing coach.
func (r *Resolver) UpdateCoachMutation(ctx context.Context, input UpdateCoachInput) (*Coach, error) {
	return r.coachService.UpdateCoach(ctx, input.CoachID, input.Name, input.Specialty)
}

// DeleteCoachMutation is a GraphQL mutation to delete a coach by ID.
func (r *Resolver) DeleteCoachMutation(ctx context.Context, coachID string) (string, error) {
	err := r.coachService.DeleteCoach(ctx, coachID)
	if err != nil {
		return "", err
	}
	return "Coach deleted successfully", nil
}

// CreateCoachInput is the input structure for creating a new coach.
type CreateCoachInput struct {
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
}

// UpdateCoachInput is the input structure for updating an existing coach.
type UpdateCoachInput struct {
	CoachID   string `json:"coachId"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
}
