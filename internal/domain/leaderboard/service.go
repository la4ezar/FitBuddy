package leaderboard

import (
	"context"
	"errors"
)

// Service handles business logic related to leaderboard entries.
type Service struct {
	repository *Repository
}

// NewService creates a new LeaderboardService instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// Create creates a new leaderboard record.
func (s *Service) Create(ctx context.Context, email string, score int) error {
	if email == "" || score < 0 {
		return errors.New("email and positive score are required")
	}

	leaderboardUser := New(email, score)

	if err := s.repository.Create(ctx, leaderboardUser); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetLeaderboardUsers(ctx context.Context) ([]*LeaderboardUser, error) {
	return s.repository.GetLeaderboardUsers(ctx)
}

// AddScore adds a score to the leaderboard for a specific user.
func (s *Service) AddScore(ctx context.Context, userID string, score float64) error {
	if userID == "" || score < 0 {
		return errors.New("user ID and a non-negative score are required")
	}

	return s.repository.AddScore(ctx, userID, score)
}

// GetTopScores retrieves the top N leaderboard entries.
func (s *Service) GetTopScores(ctx context.Context, limit int) ([]*LeaderboardUser, error) {
	if limit <= 0 {
		return nil, errors.New("limit must be a positive number")
	}

	return s.repository.GetTopScores(ctx, limit)
}
