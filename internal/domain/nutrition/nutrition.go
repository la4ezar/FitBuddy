package nutrition

import (
	"github.com/google/uuid"
	"time"
)

// Log represents a nutrition log entry in the application.
type Log struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	Meal        string    `json:"meal"`
	Description string    `json:"description"`
	Calories    int       `json:"calories"`
	LoggedAt    time.Time `json:"loggedAt"`
}

// NewLog creates a new NutritionLog instance.
func NewLog(userID, meal, description string, calories int, loggedAt time.Time) *Log {
	return &Log{
		ID:          uuid.New().String(),
		UserID:      userID,
		Meal:        meal,
		Description: description,
		Calories:    calories,
		LoggedAt:    loggedAt,
	}
}
