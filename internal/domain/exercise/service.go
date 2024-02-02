package exercise

import (
	"context"
	"errors"
)

// Service handles business logic related to exercise operations.
type Service struct {
	exerciseRepository *Repository
}

// NewService creates a new Service instance.
func NewService(exerciseRepository *Repository) *Service {
	return &Service{
		exerciseRepository: exerciseRepository,
	}
}

// CreateExercise creates a new exercise.
func (s *Service) CreateExercise(ctx context.Context, name, description string) (*Exercise, error) {
	if name == "" || description == "" {
		return nil, errors.New("name and description are required")
	}

	newExercise := NewExercise(name, description)

	if err := s.exerciseRepository.CreateExercise(ctx, newExercise); err != nil {
		return nil, err
	}

	return newExercise, nil
}

// GetExerciseByID retrieves an exercise by ID.
func (s *Service) GetExerciseByID(ctx context.Context, exerciseID string) (*Exercise, error) {
	return s.exerciseRepository.GetExerciseByID(ctx, exerciseID)
}

// UpdateExercise updates an existing exercise.
func (s *Service) UpdateExercise(ctx context.Context, exerciseID, name, description string) (*Exercise, error) {
	if name == "" || description == "" {
		return nil, errors.New("name and description are required")
	}

	existingExercise, err := s.exerciseRepository.GetExerciseByID(ctx, exerciseID)
	if err != nil {
		return nil, err
	}
	if existingExercise == nil {
		return nil, errors.New("exercise not found")
	}

	existingExercise.Name = name
	existingExercise.Description = description

	if err := s.exerciseRepository.UpdateExercise(ctx, existingExercise); err != nil {
		return nil, err
	}

	return existingExercise, nil
}

// DeleteExercise deletes an exercise by ID.
func (s *Service) DeleteExercise(ctx context.Context, exerciseID string) error {
	existingExercise, err := s.exerciseRepository.GetExerciseByID(ctx, exerciseID)
	if err != nil {
		return err
	}
	if existingExercise == nil {
		return errors.New("exercise not found")
	}

	return s.exerciseRepository.DeleteExercise(ctx, exerciseID)
}
