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

// NewResolver creates a new Resolver instance.
func NewResolver(service *Service) *Resolver {
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

// DeleteNutrition is a GraphQL mutation to delete a nutrition entry.
func (r *Resolver) DeleteNutrition(ctx context.Context, nutritionID string) (bool, error) {
	log.C(ctx).Infof("Deleting nutrition with ID %q...", nutritionID)

	err := r.service.DeleteNutrition(ctx, nutritionID)
	if err != nil {
		return false, err
	}

	log.C(ctx).Infof("Successfully deleted nutrition with ID %q", nutritionID)
	return true, nil
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
