package nutrition

import (
	"github.com/google/uuid"
	"time"
)

// Nutrition represents a nutrition entry in the application.
type Nutrition struct {
	ID        string    `json:"id"`
	UserEmail string    `json:"userEmail"`
	MealName  string    `json:"mealName"`
	Grams     int       `json:"grams"`
	Calories  int       `json:"calories"`
	CreatedAt time.Time `json:"createdAt"`
}

// NewNutrition creates a new Nutrition instance.
func NewNutrition(userEmail, mealName string, grams int, createdAt time.Time) *Nutrition {
	return &Nutrition{
		ID:        uuid.New().String(),
		UserEmail: userEmail,
		MealName:  mealName,
		Grams:     grams,
		CreatedAt: createdAt,
	}
}
