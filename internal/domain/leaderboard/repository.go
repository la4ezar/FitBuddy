package leaderboard

import (
	"context"
	"database/sql"
	"github.com/FitBuddy/pkg/log"
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

// Create creates a new leaderboard record in the database.
func (r *Repository) Create(ctx context.Context, l *LeaderboardUser) error {
	query := `
		INSERT INTO leaderboard (id, user_id, score)
		VALUES ($1, (SELECT id FROM users WHERE email = $2), $3)
	`

	_, err := r.db.ExecContext(ctx, query, l.ID, l.UserEmail, l.Score)
	return err
}

// GetLeaderboardUsers fetches the leaderboard content from the database.
func (r *Repository) GetLeaderboardUsers(ctx context.Context) ([]*LeaderboardUser, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT l.id, u.email, l.score
        FROM leaderboard l
        JOIN users u ON l.user_id = u.id                                
        ORDER BY l.score DESC
    `)
	if err != nil {
		log.C(ctx).Info(err)
		return nil, err
	}
	defer rows.Close()

	var leaderboardUsers []*LeaderboardUser

	for rows.Next() {
		var leaderboardUser LeaderboardUser
		if err := rows.Scan(&leaderboardUser.ID, &leaderboardUser.UserEmail, &leaderboardUser.Score); err != nil {
			return nil, err
		}
		leaderboardUsers = append(leaderboardUsers, &leaderboardUser)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return leaderboardUsers, nil
}

// AddScore adds a score to the leaderboard for a specific user.
func (r *Repository) AddScore(ctx context.Context, email string, score float64) error {
	query := `
		UPDATE leaderboard
		SET score = leaderboard.score + $2
		WHERE user_id = (SELECT id FROM users WHERE email = $1)
	`

	_, err := r.db.ExecContext(ctx, query, email, score)
	return err
}
