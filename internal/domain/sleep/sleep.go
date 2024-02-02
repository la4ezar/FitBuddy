package sleep

import (
	"github.com/google/uuid"
	"time"
)

// Log represents a sleep log entry in the application.
type Log struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Duration  int       `json:"duration"`
	SleepTime time.Time `json:"sleepTime"`
	WakeTime  time.Time `json:"wakeTime"`
}

// NewLog creates a new SleepLog instance.
func NewLog(userID string, duration int, sleepTime, wakeTime time.Time) *Log {
	return &Log{
		ID:        uuid.New().String(),
		UserID:    userID,
		Duration:  duration,
		SleepTime: sleepTime,
		WakeTime:  wakeTime,
	}
}
