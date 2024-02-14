package nutrition

import (
	"context"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
	"time"
)

// Resolver handles GraphQL queries and mutations for the Log aggregate.
type Resolver struct {
	service *Service
}

// NewNutritionResolver creates a new NutritionResolver instance.
func NewNutritionResolver(service *Service) *Resolver {
	return &Resolver{
		service: service,
	}
}

// CreateNutrition is a GraphQL mutation to create a new nutrition entry.
func (r *Resolver) CreateNutrition(ctx context.Context, email, meal, date string, servingSize, numberOfServings int) (*graphql.Nutrition, error) {
	log.C(ctx).Infof("Creating Nutrition for user with email %q...", email)

	createdAt, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, err
	}

	nutrition, err := r.service.CreateNutrition(ctx, email, meal, servingSize, numberOfServings, createdAt)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully created nutrition for user with email %q", email)

	gqlNutrition := &graphql.Nutrition{
		ID:        nutrition.ID,
		UserEmail: nutrition.UserEmail,
		MealName:  nutrition.MealName,
		Grams:     nutrition.Grams,
		Date:      nutrition.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return gqlNutrition, nil
}

// GetAllNutritions is a GraphQL query to retrieve all nutritions.
func (r *Resolver) GetAllNutritions(ctx context.Context, email, date string) ([]*graphql.Nutrition, error) {
	log.C(ctx).Info("Getting all nutritions...")
	createdAt, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, err
	}
	nutritions, err := r.service.GetAllNutritions(ctx, email, createdAt)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Info("Successfully got all nutritions")

	gqlNutritions := make([]*graphql.Nutrition, 0, len(nutritions))
	for _, n := range nutritions {
		gqlNutritions = append(gqlNutritions, &graphql.Nutrition{
			ID:        n.ID,
			UserEmail: n.UserEmail,
			MealName:  n.MealName,
			Grams:     n.Grams,
			Calories:  n.Calories,
			Date:      n.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return gqlNutritions, nil
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
