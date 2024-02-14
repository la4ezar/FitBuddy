package sleep

import (
	"context"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
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

func (r *Resolver) CreateSleepLog(ctx context.Context, userEmail string, sleepTime, wakeTime, date time.Time) (*graphql.SleepLog, error) {
	sleepTimeParsed := sleepTime.Format("2006-01-02 15:04:05")
	wakeTimeParsed := wakeTime.Format("2006-01-02 15:04:05")
	dateParsed := date.Format("2006-01-02 15:04:05")
	log.C(ctx).Infof("Creating sleep with sleep time %q and wake time %q for day %q for user with email %q...", sleepTimeParsed, wakeTimeParsed, dateParsed, userEmail)

	sleep, err := r.service.CreateLog(ctx, userEmail, sleepTime, wakeTime, date)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully created sleep with sleep time %q and wake time %q for day %q for user with email %q", sleepTimeParsed, wakeTimeParsed, dateParsed, userEmail)

	return &graphql.SleepLog{
		ID:        sleep.ID,
		SleepTime: sleepTimeParsed,
		WakeTime:  wakeTimeParsed,
	}, nil
}

func (r *Resolver) GetSleepLogByEmailAndDate(ctx context.Context, userEmail string, date time.Time) ([]*graphql.SleepLog, error) {
	dateParsed := date.Format("2006-01-02 15:04:05")
	log.C(ctx).Infof("Getting sleep for date %q and user with email %q...", dateParsed, userEmail)

	sleep, err := r.service.GetSleepLogByEmailAndDate(ctx, userEmail, date)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully getting sleep for date %q and user with email %q", dateParsed, userEmail)

	gqlSleep := make([]*graphql.SleepLog, 0, len(sleep))
	for _, s := range sleep {
		gqlSleep = append(gqlSleep, &graphql.SleepLog{
			ID:        s.ID,
			SleepTime: s.SleepTime.Format("2006-01-02 15:04:05"),
			WakeTime:  s.WakeTime.Format("2006-01-02 15:04:05"),
		})
	}

	return gqlSleep, nil
}

func (r *Resolver) DeleteSleepLog(ctx context.Context, sleepID string) (bool, error) {
	log.C(ctx).Infof("Deleting sleep with ID %q...", sleepID)

	err := r.service.DeleteLog(ctx, sleepID)
	if err != nil {
		return false, err
	}
	log.C(ctx).Infof("Successfully deleted sleep with ID %q", sleepID)

	return true, nil
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
