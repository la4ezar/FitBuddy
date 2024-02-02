package leaderboard

import (
	"context"
)

// Resolver handles GraphQL queries and mutations for the Leaderboard aggregate.
type Resolver struct {
	leaderboardService *Service
}

// NewLeaderboardResolver creates a new LeaderboardResolver instance.
func NewLeaderboardResolver(leaderboardService *Service) *Resolver {
	return &Resolver{
		leaderboardService: leaderboardService,
	}
}

// AddScoreMutation is a GraphQL mutation to add a score to the leaderboard for a specific user.
func (r *Resolver) AddScoreMutation(ctx context.Context, input AddScoreInput) (string, error) {
	err := r.leaderboardService.AddScore(ctx, input.UserID, input.Score)
	if err != nil {
		return "", err
	}
	return "Score added to the leaderboard successfully", nil
}

// GetTopScoresQuery is a GraphQL query to retrieve the top N leaderboard entries.
func (r *Resolver) GetTopScoresQuery(ctx context.Context, limit int) ([]*Leaderboard, error) {
	return r.leaderboardService.GetTopScores(ctx, limit)
}

// AddScoreInput is the input structure for adding a score to the leaderboard.
type AddScoreInput struct {
	UserID string  `json:"userId"`
	Score  float64 `json:"score"`
}
