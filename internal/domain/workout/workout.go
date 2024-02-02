package workout

import (
	"github.com/google/uuid"
	"time"
)

// Log represents a workout log entry in the application.
type Log struct {
	ID       string    `json:"id"`
	UserID   string    `json:"userId"`
	Exercise string    `json:"exercise"`
	Sets     int       `json:"sets"`
	Reps     int       `json:"reps"`
	Weight   float64   `json:"weight"`
	LoggedAt time.Time `json:"loggedAt"`
}

// NewLog creates a new WorkoutLog instance.
func NewLog(userID, exercise string, sets, reps int, weight float64, loggedAt time.Time) *Log {
	return &Log{
		ID:       uuid.New().String(),
		UserID:   userID,
		Exercise: exercise,
		Sets:     sets,
		Reps:     reps,
		Weight:   weight,
		LoggedAt: loggedAt,
	}
}
