package leaderboard

import (
	"context"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
)

// Resolver handles GraphQL queries and mutations for the Leaderboard aggregate.
type Resolver struct {
	service *Service
}

// NewLeaderboardResolver creates a new LeaderboardResolver instance.
func NewLeaderboardResolver(leaderboardService *Service) *Resolver {
	return &Resolver{
		service: leaderboardService,
	}
}

// GetLeaderboardUsers is a GraphQL query to retrieve leaderboard content.
func (r *Resolver) GetLeaderboardUsers(ctx context.Context) ([]*graphql.LeaderboardUser, error) {
	log.C(ctx).Info("Getting leaderboard...")

	leaderboardUsers, err := r.service.GetLeaderboardUsers(ctx)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Info("Successfully got leaderboard")

	gqlLeaderboardUsers := make([]*graphql.LeaderboardUser, 0, len(leaderboardUsers))
	for _, l := range leaderboardUsers {
		gqlLeaderboardUsers = append(gqlLeaderboardUsers, &graphql.LeaderboardUser{
			ID:        l.ID,
			UserEmail: l.UserEmail,
			Score:     l.Score,
		})
	}

	return gqlLeaderboardUsers, nil
}

// AddScoreMutation is a GraphQL mutation to add a score to the leaderboard for a specific user.
func (r *Resolver) AddScoreMutation(ctx context.Context, input AddScoreInput) (string, error) {
	err := r.service.AddScore(ctx, input.UserID, input.Score)
	if err != nil {
		return "", err
	}
	return "Score added to the leaderboard successfully", nil
}

// GetTopScoresQuery is a GraphQL query to retrieve the top N leaderboard entries.
func (r *Resolver) GetTopScoresQuery(ctx context.Context, limit int) ([]*LeaderboardUser, error) {
	return r.service.GetTopScores(ctx, limit)
}

// AddScoreInput is the input structure for adding a score to the leaderboard.
type AddScoreInput struct {
	UserID string  `json:"userId"`
	Score  float64 `json:"score"`
}
