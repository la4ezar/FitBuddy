package meal

import (
	"context"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
)

// Resolver handles GraphQL queries and mutations for the Meal aggregate.
type Resolver struct {
	service *Service
}

// NewResolver creates a new Meal Resolver instance.
func NewResolver(service *Service) *Resolver {
	return &Resolver{
		service: service,
	}
}

// GetAllMeals is a GraphQL query to retrieve all meals.
func (r *Resolver) GetAllMeals(ctx context.Context) ([]*graphql.Meal, error) {
	log.C(ctx).Info("Getting all meals...")

	meals, err := r.service.GetAllMeals(ctx)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Info("Successfully got all meals")

	gqlMeals := make([]*graphql.Meal, 0, len(meals))
	for _, m := range meals {
		gqlMeals = append(gqlMeals, &graphql.Meal{
			ID:   m.ID,
			Name: m.Name,
		})
	}

	return gqlMeals, nil
}
