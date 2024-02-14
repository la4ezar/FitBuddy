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
