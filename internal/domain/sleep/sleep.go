package sleep

import (
	"github.com/google/uuid"
	"time"
)

// Log represents a sleep log entry in the application.
type Log struct {
	ID        string    `json:"id"`
	UserEmail string    `json:"userEmail"`
	SleepTime time.Time `json:"sleepTime"`
	WakeTime  time.Time `json:"wakeTime"`
	LoggedAt  time.Time `json:"loggedAt"`
}

// NewLog creates a new SleepLog instance.
func NewLog(userEmail string, sleepTime, wakeTime, loggedAt time.Time) *Log {
	return &Log{
		ID:        uuid.New().String(),
		UserEmail: userEmail,
		SleepTime: sleepTime,
		WakeTime:  wakeTime,
		LoggedAt:  loggedAt,
	}
}
