package goal

import (
	"github.com/google/uuid"
	"time"
)

// Goal represents a fitness or wellness goal in the application.
type Goal struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}

// NewGoal creates a new Goal instance.
func NewGoal(userID, title, description string, startDate, endDate time.Time) *Goal {
	return &Goal{
		ID:          uuid.New().String(),
		UserID:      userID,
		Title:       title,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
	}
}
