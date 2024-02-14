package sleep

import (
	"context"
	"errors"
	"time"
)

// Service handles business logic related to sleep log entries.
type Service struct {
	repository *Repository
}

// NewService creates a new Service instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// CreateLog creates a new sleep log entry.
func (s *Service) CreateLog(ctx context.Context, userEmail string, sleepTime, wakeTime, date time.Time) (*Log, error) {
	if userEmail == "" {
		return nil, errors.New("user email and valid sleep and wake times are required")
	}

	if wakeTime.Before(sleepTime) {
		sleepTime = sleepTime.AddDate(0, 0, -1)
	}
	newLog := NewLog(userEmail, sleepTime, wakeTime, date)

	if err := s.repository.CreateLog(ctx, newLog); err != nil {
		return nil, err
	}

	return newLog, nil
}

// GetSleepLogByEmailAndDate gets sleep for a given date and user by email
func (s *Service) GetSleepLogByEmailAndDate(ctx context.Context, userEmail string, date time.Time) ([]*Log, error) {
	if userEmail == "" {
		return nil, errors.New("user email is required")
	}

	sleep, err := s.repository.GetSleepLogByEmailAndDate(ctx, userEmail, date)
	if err != nil {
		return nil, err
	}

	return sleep, nil
}

// DeleteLog deletes a sleep log entry by ID.
func (s *Service) DeleteLog(ctx context.Context, sleepLogID string) error {
	existingLog, err := s.repository.GetLogByID(ctx, sleepLogID)
	if err != nil {
		return err
	}
	if existingLog == nil {
		return errors.New("sleep log entry not found")
	}

	return s.repository.DeleteLog(ctx, sleepLogID)
}
