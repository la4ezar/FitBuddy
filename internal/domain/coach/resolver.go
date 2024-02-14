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

// UnbookCoach is an unbooking a coach with name
func (r *Resolver) UnbookCoach(ctx context.Context, email, coachName string) (bool, error) {
	log.C(ctx).Infof("Unbooking a coach with name %q for user with email %q...", coachName, email)

	if _, err := r.service.UnbookCoach(ctx, email, coachName); err != nil {
		return false, err
	}
	log.C(ctx).Infof("Successfully unbooked a coach with name %q for user with email %q", coachName, email)

	return true, nil
}
