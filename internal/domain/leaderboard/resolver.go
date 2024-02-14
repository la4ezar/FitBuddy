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
