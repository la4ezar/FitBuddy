package sleep

import (
	"context"
	"time"
)

// Resolver handles GraphQL queries and mutations for the Log aggregate.
type Resolver struct {
	service *Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(service *Service) *Resolver {
	return &Resolver{
		service: service,
	}
}

// CreateLogMutation is a GraphQL mutation to create a new sleep log entry.
func (r *Resolver) CreateLogMutation(ctx context.Context, input CreateLogInput) (*Log, error) {
	return r.service.CreateLog(ctx, input.UserID, input.Duration, input.SleepTime, input.WakeTime)
}

// GetLogQuery is a GraphQL query to retrieve a sleep log entry by ID.
func (r *Resolver) GetLogQuery(ctx context.Context, sleepLogID string) (*Log, error) {
	return r.service.GetLogByID(ctx, sleepLogID)
}

// UpdateLogMutation is a GraphQL mutation to update an existing sleep log entry.
func (r *Resolver) UpdateLogMutation(ctx context.Context, input UpdateLogInput) (*Log, error) {
	return r.service.UpdateLog(ctx, input.LogID, input.Duration, input.SleepTime, input.WakeTime)
}

// DeleteLogMutation is a GraphQL mutation to delete a sleep log entry by ID.
func (r *Resolver) DeleteLogMutation(ctx context.Context, sleepLogID string) (string, error) {
	err := r.service.DeleteLog(ctx, sleepLogID)
	if err != nil {
		return "", err
	}
	return "Sleep log entry deleted successfully", nil
}

// CreateLogInput is the input structure for creating a new sleep log entry.
type CreateLogInput struct {
	UserID    string    `json:"userId"`
	Duration  int       `json:"duration"`
	SleepTime time.Time `json:"sleepTime"`
	WakeTime  time.Time `json:"wakeTime"`
}

// UpdateLogInput is the input structure for updating an existing sleep log entry.
type UpdateLogInput struct {
	LogID     string    `json:"sleepLogId"`
	Duration  int       `json:"duration"`
	SleepTime time.Time `json:"sleepTime"`
	WakeTime  time.Time `json:"wakeTime"`
}
