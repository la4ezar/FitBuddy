package exercise

import (
	"context"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
)

// Resolver handles GraphQL queries and mutations for the Exercise aggregate.
type Resolver struct {
	service *Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(service *Service) *Resolver {
	return &Resolver{
		service: service,
	}
}

// GetAllExercises is a GraphQL query to retrieve all exercises.
func (r *Resolver) GetAllExercises(ctx context.Context) ([]*graphql.Exercise, error) {
	log.C(ctx).Info("Getting all exercises...")

	exercises, err := r.service.GetAllExercises(ctx)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Info("Successfully got all exercises")

	gqlExercises := make([]*graphql.Exercise, 0, len(exercises))
	for _, e := range exercises {
		gqlExercises = append(gqlExercises, &graphql.Exercise{
			ID:   e.ID,
			Name: e.Name,
		})
	}

	return gqlExercises, nil
}
