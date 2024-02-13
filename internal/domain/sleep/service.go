package sleep

import (
	"context"
	"errors"
	"time"
)

// Service handles business logic related to sleep log entries.
type Service struct {
	sleepRepository *Repository
}

// NewService creates a new Service instance.
func NewService(sleepRepository *Repository) *Service {
	return &Service{
		sleepRepository: sleepRepository,
	}
}

// CreateLog creates a new sleep log entry.
func (s *Service) CreateLog(ctx context.Context, userEmail string, sleepTime, wakeTime, date time.Time) (*Log, error) {
	if userEmail == "" {
		return nil, errors.New("user email and valid sleep and wake times are required")
	}

	if wakeTime.Before(sleepTime) {
		sleepTime.AddDate(0, 0, -1)
	}

	newLog := NewLog(userEmail, sleepTime, wakeTime, date)

	if err := s.sleepRepository.CreateLog(ctx, newLog); err != nil {
		return nil, err
	}

	return newLog, nil
}

// GetSleepByEmailAndDate gets sleep for a given date and user by email
func (s *Service) GetSleepByEmailAndDate(ctx context.Context, userEmail string, date time.Time) ([]*Log, error) {
	if userEmail == "" {
		return nil, errors.New("user email is required")
	}

	sleep, err := s.sleepRepository.GetSleepByEmailAndDate(ctx, userEmail, date)
	if err != nil {
		return nil, err
	}

	return sleep, nil
}

// GetLogByID retrieves a sleep log entry by ID.
func (s *Service) GetLogByID(ctx context.Context, sleepLogID string) (*Log, error) {
	return s.sleepRepository.GetLogByID(ctx, sleepLogID)
}

// UpdateLog updates an existing sleep log entry.
func (s *Service) UpdateLog(ctx context.Context, sleepLogID string, duration int, sleepTime, wakeTime time.Time) (*Log, error) {
	if duration <= 0 || wakeTime.Before(sleepTime) {
		return nil, errors.New("positive duration and valid sleep and wake times are required")
	}

	existingLog, err := s.sleepRepository.GetLogByID(ctx, sleepLogID)
	if err != nil {
		return nil, err
	}
	if existingLog == nil {
		return nil, errors.New("sleep log entry not found")
	}

	existingLog.SleepTime = sleepTime
	existingLog.WakeTime = wakeTime

	if err := s.sleepRepository.UpdateLog(ctx, existingLog); err != nil {
		return nil, err
	}

	return existingLog, nil
}

// DeleteLog deletes a sleep log entry by ID.
func (s *Service) DeleteLog(ctx context.Context, sleepLogID string) error {
	existingLog, err := s.sleepRepository.GetLogByID(ctx, sleepLogID)
	if err != nil {
		return err
	}
	if existingLog == nil {
		return errors.New("sleep log entry not found")
	}

	return s.sleepRepository.DeleteLog(ctx, sleepLogID)
}
