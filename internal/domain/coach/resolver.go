package coach

import (
	"context"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
)

// Resolver handles GraphQL queries and mutations for the Coach aggregate.
type Resolver struct {
	service *Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(service *Service) *Resolver {
	return &Resolver{
		service: service,
	}
}

// CreateCoachMutation is a GraphQL mutation to create a new coach.
func (r *Resolver) CreateCoachMutation(ctx context.Context, input CreateCoachInput) (*Coach, error) {
	return r.service.CreateCoach(ctx, input.Name, input.Specialty)
}

// GetCoachQuery is a GraphQL query to retrieve a coach by ID.
func (r *Resolver) GetCoachQuery(ctx context.Context, coachID string) (*Coach, error) {
	return r.service.GetCoachByID(ctx, coachID)
}

// GetAllCoaches is a GraphQL query to retrieve all coaches.
func (r *Resolver) GetAllCoaches(ctx context.Context) ([]*graphql.Coach, error) {
	log.C(ctx).Info("Getting all coaches...")

	coaches, err := r.service.GetAllCoaches(ctx)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Info("Successfully got all coaches")

	gqlCoaches := make([]*graphql.Coach, 0, len(coaches))
	for _, c := range coaches {
		gqlCoaches = append(gqlCoaches, &graphql.Coach{
			ID:        c.ID,
			ImageURL:  c.ImageURL,
			Name:      c.Name,
			Specialty: c.Specialty,
		})
	}

	return gqlCoaches, nil
}

// IsCoachBookedByUser is a GraphQL query to tell if coach is booked by a given user.
func (r *Resolver) IsCoachBookedByUser(ctx context.Context, coachName, userEmail string) (bool, error) {
	log.C(ctx).Infof("Checking if coach with name %q is booked by user with email %q...", coachName, userEmail)

	isBooked, err := r.service.IsCoachBookedByUser(ctx, coachName, userEmail)
	if err != nil {
		return false, err
	}
	log.C(ctx).Infof("Successfully checked if coach with name %q is booked by user with email %q. The response is: %v", coachName, userEmail, isBooked)

	return isBooked, nil
}

// IsCoachBooked is a GraphQL query to tell if coach is booked by a some user.
func (r *Resolver) IsCoachBooked(ctx context.Context, coachName string) (bool, error) {
	log.C(ctx).Infof("Checking if coach with name %q is booked...", coachName)

	isBooked, err := r.service.IsCoachBooked(ctx, coachName)
	if err != nil {
		return false, err
	}
	log.C(ctx).Infof("Successfully checked if coach with name %q is booked. The response is: %v", coachName, isBooked)

	return isBooked, nil
}

// BookCoach is a booking a coach with name
func (r *Resolver) BookCoach(ctx context.Context, email, coachName string) (bool, error) {
	log.C(ctx).Infof("Booking a coach with name %q for user with email %q...", coachName, email)

	if _, err := r.service.BookCoach(ctx, email, coachName); err != nil {
		return false, err
	}
	log.C(ctx).Infof("Successfully booked a coach with name %q for user with email %q", coachName, email)

	return true, nil
}

// UnbookCoach is a booking a coach with name
func (r *Resolver) UnbookCoach(ctx context.Context, email, coachName string) (bool, error) {
	log.C(ctx).Infof("Unbooking a coach with name %q for user with email %q...", coachName, email)

	if _, err := r.service.UnbookCoach(ctx, email, coachName); err != nil {
		return false, err
	}
	log.C(ctx).Infof("Successfully unbooked a coach with name %q for user with email %q", coachName, email)

	return true, nil
}

// UpdateCoachMutation is a GraphQL mutation to update an existing coach.
func (r *Resolver) UpdateCoachMutation(ctx context.Context, input UpdateCoachInput) (*Coach, error) {
	return r.service.UpdateCoach(ctx, input.CoachID, input.Name, input.Specialty)
}

// DeleteCoachMutation is a GraphQL mutation to delete a coach by ID.
func (r *Resolver) DeleteCoachMutation(ctx context.Context, coachID string) (string, error) {
	err := r.service.DeleteCoach(ctx, coachID)
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
