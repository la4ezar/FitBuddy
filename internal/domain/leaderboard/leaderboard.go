package leaderboard

// Leaderboard represents a leaderboard entry in the application.
type Leaderboard struct {
	UserID   string  `json:"userId"`
	Username string  `json:"username"`
	Score    float64 `json:"score"`
}

// NewLeaderboard creates a new Leaderboard instance.
func NewLeaderboard(userID, username string, score float64) *Leaderboard {
	return &Leaderboard{
		UserID:   userID,
		Username: username,
		Score:    score,
	}
}
