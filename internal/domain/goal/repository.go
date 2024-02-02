package goal

import (
	"context"
	"database/sql"
)

// Repository is a repository for managing fitness and wellness goals data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreateGoal creates a new fitness or wellness goal in the database.
func (r *Repository) CreateGoal(ctx context.Context, goal *Goal) error {
	query := `
		INSERT INTO goals (id, user_id, title, description, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		goal.ID,
		goal.UserID,
		goal.Title,
		goal.Description,
		goal.StartDate,
		goal.EndDate,
	)

	return err
}

// GetGoalByID retrieves a fitness or wellness goal from the database by ID.
func (r *Repository) GetGoalByID(ctx context.Context, goalID string) (*Goal, error) {
	query := `
		SELECT id, user_id, title, description, start_date, end_date
		FROM goals
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, goalID)

	var goalItem Goal
	err := row.Scan(
		&goalItem.ID,
		&goalItem.UserID,
		&goalItem.Title,
		&goalItem.Description,
		&goalItem.StartDate,
		&goalItem.EndDate,
	)
	if err == sql.ErrNoRows {
		return nil, nil // Goal not found
	} else if err != nil {
		return nil, err
	}

	return &goalItem, nil
}

// UpdateGoal updates an existing fitness or wellness goal in the database.
func (r *Repository) UpdateGoal(ctx context.Context, goal *Goal) error {
	query := `
		UPDATE goals
		SET title = $2, description = $3, start_date = $4, end_date = $5
		WHERE id = $1
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		goal.ID,
		goal.Title,
		goal.Description,
		goal.StartDate,
		goal.EndDate,
	)

	return err
}

// DeleteGoal deletes a fitness or wellness goal from the database by ID.
func (r *Repository) DeleteGoal(ctx context.Context, goalID string) error {
	query := `
		DELETE FROM goals
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, goalID)
	return err
}
