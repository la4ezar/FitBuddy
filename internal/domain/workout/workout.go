package workout

import (
	"github.com/google/uuid"
	"time"
)

// Workout represents a workout entry in the application.
type Workout struct {
	ID           string    `json:"id"`
	UserEmail    string    `json:"userEmail"`
	ExerciseName string    `json:"exerciseName"`
	Sets         int       `json:"sets"`
	Reps         int       `json:"reps"`
	Weight       float64   `json:"weight"`
	CreatedAt    time.Time `json:"createdAt"`
}

// NewWorkout creates a new Workout instance.
func NewWorkout(userID, exercise string, sets, reps int, weight float64, createdAt time.Time) *Workout {
	return &Workout{
		ID:           uuid.New().String(),
		UserEmail:    userID,
		ExerciseName: exercise,
		Sets:         sets,
		Reps:         reps,
		Weight:       weight,
		CreatedAt:    createdAt,
	}
}
