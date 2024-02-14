package nutrition

import (
	"context"
	"errors"
	"github.com/FitBuddy/pkg/log"
	"time"
)

// Service handles business logic related to nutrition log entries.
type Service struct {
	repository *Repository
}

// NewService creates a new NutritionService instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// CreateNutrition creates a new nutrition entry.
func (s *Service) CreateNutrition(ctx context.Context, email, meal string, servingSize, numberOfServings int, createdAt time.Time) (*Nutrition, error) {
	if email == "" || meal == "" || servingSize <= 0 || numberOfServings <= 0 || createdAt == (time.Time{}) {
		return nil, errors.New("email, meal, createdAt time, positive servingSize and positive numberOfServings are required")
	}

	nutrition := NewNutrition(email, meal, servingSize*numberOfServings, createdAt)

	if err := s.repository.CreateNutrition(ctx, nutrition); err != nil {
		log.C(ctx).Infof("Creating Nutrition for user with email %q...", err)
		return nil, err
	}

	return nutrition, nil
}

// DeleteNutrition deletes a nutrition entry.
func (s *Service) DeleteNutrition(ctx context.Context, nutritionID string) error {
	if nutritionID == "" {
		return errors.New("nutritionID is required")
	}

	if err := s.repository.DeleteNutrition(ctx, nutritionID); err != nil {
		log.C(ctx).Infof("Deleting nutrition with ID %q is returning %v", nutritionID, err)
		return err
	}

	return nil
}

// GetAllNutritions retrieves all nutritions.
func (s *Service) GetAllNutritions(ctx context.Context, email string, date time.Time) ([]*Nutrition, error) {
	return s.repository.GetAllNutritions(ctx, email, date)
}
