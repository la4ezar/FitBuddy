package leaderboard

import "github.com/google/uuid"

// LeaderboardUser represents a leaderboard entry in the application.
type LeaderboardUser struct {
	ID        string `json:"id"`
	UserEmail string `json:"userEmail"`
	Score     int    `json:"score"`
}

// New creates a new Leaderboard instance.
func New(userEmail string, score int) *LeaderboardUser {
	return &LeaderboardUser{
		ID:        uuid.New().String(),
		UserEmail: userEmail,
		Score:     score,
	}
}
