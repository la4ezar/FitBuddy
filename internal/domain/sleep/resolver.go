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

// CreateSleepLog creates a sleep log
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

// GetSleepLogByEmailAndDate gets sleep lig by user email and date
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

// DeleteSleepLog deletes sleep log
func (r *Resolver) DeleteSleepLog(ctx context.Context, sleepID string) (bool, error) {
	log.C(ctx).Infof("Deleting sleep with ID %q...", sleepID)

	err := r.service.DeleteLog(ctx, sleepID)
	if err != nil {
		return false, err
	}
	log.C(ctx).Infof("Successfully deleted sleep with ID %q", sleepID)

	return true, nil
}
