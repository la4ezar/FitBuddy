package leaderboard

import (
	"context"
	"errors"
)

// Service handles business logic related to leaderboard entries.
type Service struct {
	leaderboardRepository *Repository
}

// NewService creates a new LeaderboardService instance.
func NewService(leaderboardRepository *Repository) *Service {
	return &Service{
		leaderboardRepository: leaderboardRepository,
	}
}

// AddScore adds a score to the leaderboard for a specific user.
func (s *Service) AddScore(ctx context.Context, userID string, score float64) error {
	if userID == "" || score < 0 {
		return errors.New("user ID and a non-negative score are required")
	}

	return s.leaderboardRepository.AddScore(ctx, userID, score)
}

// GetTopScores retrieves the top N leaderboard entries.
func (s *Service) GetTopScores(ctx context.Context, limit int) ([]*Leaderboard, error) {
	if limit <= 0 {
		return nil, errors.New("limit must be a positive number")
	}

	return s.leaderboardRepository.GetTopScores(ctx, limit)
}
