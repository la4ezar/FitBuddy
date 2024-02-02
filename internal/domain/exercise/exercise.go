package exercise

import "github.com/google/uuid"

// Exercise represents a fitness exercise in the application.
type Exercise struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// NewExercise creates a new Exercise instance.
func NewExercise(name, description string) *Exercise {
	return &Exercise{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
	}
}
