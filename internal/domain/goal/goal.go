package goal

import (
	"github.com/google/uuid"
	"time"
)

// Goal represents a fitness goal in the application.
type Goal struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Completed   bool      `json:"completed"`
}

// NewGoal creates a new Goal instance.
func NewGoal(name, description string, startDate, endDate time.Time) *Goal {
	return &Goal{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
		Completed:   false,
	}
}
