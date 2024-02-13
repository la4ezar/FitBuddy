package exercise

import (
	"context"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
)

// Resolver handles GraphQL queries and mutations for the Exercise aggregate.
type Resolver struct {
	exerciseService *Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(exerciseService *Service) *Resolver {
	return &Resolver{
		exerciseService: exerciseService,
	}
}

// CreateExerciseMutation is a GraphQL mutation to create a new
func (r *Resolver) CreateExerciseMutation(ctx context.Context, input CreateExerciseInput) (*Exercise, error) {
	return r.exerciseService.CreateExercise(ctx, input.Name, input.Description)
}

// GetExerciseQuery is a GraphQL query to retrieve an exercise by ID.
func (r *Resolver) GetExerciseQuery(ctx context.Context, exerciseID string) (*Exercise, error) {
	return r.exerciseService.GetExerciseByID(ctx, exerciseID)
}

// GetAllExercises is a GraphQL query to retrieve all exercises.
func (r *Resolver) GetAllExercises(ctx context.Context) ([]*graphql.Exercise, error) {
	log.C(ctx).Info("Getting all exercises...")

	exercises, err := r.exerciseService.GetAllExercises(ctx)
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

// UpdateExerciseMutation is a GraphQL mutation to update an existing
func (r *Resolver) UpdateExerciseMutation(ctx context.Context, input UpdateExerciseInput) (*Exercise, error) {
	return r.exerciseService.UpdateExercise(ctx, input.ExerciseID, input.Name, input.Description)
}

// DeleteExerciseMutation is a GraphQL mutation to delete an exercise by ID.
func (r *Resolver) DeleteExerciseMutation(ctx context.Context, exerciseID string) (string, error) {
	err := r.exerciseService.DeleteExercise(ctx, exerciseID)
	if err != nil {
		return "", err
	}
	return "Exercise deleted successfully", nil
}

// CreateExerciseInput represents the input for the CreateExerciseMutation.
type CreateExerciseInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateExerciseInput represents the input for the UpdateExerciseMutation.
type UpdateExerciseInput struct {
	ExerciseID  string `json:"exerciseID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
