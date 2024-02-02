package nutrition

import (
	"context"
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

// CreateLogMutation is a GraphQL mutation to create a new nutrition log entry.
func (r *Resolver) CreateLogMutation(ctx context.Context, input CreateLogInput) (*Log, error) {
	return r.service.CreateLog(ctx, input.UserID, input.Meal, input.Description, input.Calories, input.LoggedAt)
}

// GetLogQuery is a GraphQL query to retrieve a nutrition log entry by ID.
func (r *Resolver) GetLogQuery(ctx context.Context, nutritionLogID string) (*Log, error) {
	return r.service.GetLogByID(ctx, nutritionLogID)
}

// UpdateLogMutation is a GraphQL mutation to update an existing nutrition log entry.
func (r *Resolver) UpdateLogMutation(ctx context.Context, input UpdateLogInput) (*Log, error) {
	return r.service.UpdateLog(ctx, input.LogID, input.Meal, input.Description, input.Calories, input.LoggedAt)
}

// DeleteLogMutation is a GraphQL mutation to delete a nutrition log entry by ID.
func (r *Resolver) DeleteLogMutation(ctx context.Context, nutritionLogID string) (string, error) {
	err := r.service.DeleteLog(ctx, nutritionLogID)
	if err != nil {
		return "", err
	}
	return "Nutrition log entry deleted successfully", nil
}

// CreateLogInput is the input structure for creating a new nutrition log entry.
type CreateLogInput struct {
	UserID      string    `json:"userId"`
	Meal        string    `json:"meal"`
	Description string    `json:"description"`
	Calories    int       `json:"calories"`
	LoggedAt    time.Time `json:"loggedAt"`
}

// UpdateLogInput is the input structure for updating an existing nutrition log entry.
type UpdateLogInput struct {
	LogID       string    `json:"nutritionLogId"`
	Meal        string    `json:"meal"`
	Description string    `json:"description"`
	Calories    int       `json:"calories"`
	LoggedAt    time.Time `json:"loggedAt"`
}
