package leaderboard

import (
	"context"
	"database/sql"
)

// Repository is a repository for managing leaderboard entries.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new LeaderboardRepository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// AddScore adds a score to the leaderboard for a specific user.
func (r *Repository) AddScore(ctx context.Context, userID string, score float64) error {
	query := `
		INSERT INTO leaderboard (user_id, score)
		VALUES ($1, $2)
		ON CONFLICT (user_id) DO UPDATE SET score = leaderboard.score + EXCLUDED.score
	`

	_, err := r.db.ExecContext(ctx, query, userID, score)
	return err
}

// GetTopScores retrieves the top N leaderboard entries.
func (r *Repository) GetTopScores(ctx context.Context, limit int) ([]*Leaderboard, error) {
	query := `
		SELECT user_id, username, score
		FROM leaderboard
		ORDER BY score DESC
		LIMIT $1
	`

	rows, err := r.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topScores []*Leaderboard
	for rows.Next() {
		var l Leaderboard
		err := rows.Scan(&l.UserID, &l.Username, &l.Score)
		if err != nil {
			return nil, err
		}
		topScores = append(topScores, &l)
	}

	return topScores, nil
}
