package workout

import (
	"context"
	"database/sql"
)

// Repository is a repository for managing workout log data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreateLog creates a new workout log entry in the database.
func (r *Repository) CreateLog(ctx context.Context, workoutLog *Log) error {
	query := `
		INSERT INTO workout_logs (id, user_id, exercise, sets, reps, weight, logged_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		workoutLog.ID,
		workoutLog.UserID,
		workoutLog.Exercise,
		workoutLog.Sets,
		workoutLog.Reps,
		workoutLog.Weight,
		workoutLog.LoggedAt,
	)

	return err
}

// GetLogByID retrieves a workout log entry from the database by ID.
func (r *Repository) GetLogByID(ctx context.Context, workoutLogID string) (*Log, error) {
	query := `
		SELECT id, user_id, exercise, sets, reps, weight, logged_at
		FROM workout_logs
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, workoutLogID)

	var workoutLogItem Log
	err := row.Scan(
		&workoutLogItem.ID,
		&workoutLogItem.UserID,
		&workoutLogItem.Exercise,
		&workoutLogItem.Sets,
		&workoutLogItem.Reps,
		&workoutLogItem.Weight,
		&workoutLogItem.LoggedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil // Workout log not found
	} else if err != nil {
		return nil, err
	}

	return &workoutLogItem, nil
}

// UpdateLog updates an existing workout log entry in the database.
func (r *Repository) UpdateLog(ctx context.Context, workoutLog *Log) error {
	query := `
		UPDATE workout_logs
		SET exercise = $3, sets = $4, reps = $5, weight = $6, logged_at = $7
		WHERE id = $1 AND user_id = $2
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		workoutLog.ID,
		workoutLog.UserID,
		workoutLog.Exercise,
		workoutLog.Sets,
		workoutLog.Reps,
		workoutLog.Weight,
		workoutLog.LoggedAt,
	)

	return err
}

// DeleteLog deletes a workout log entry from the database by ID.
func (r *Repository) DeleteLog(ctx context.Context, workoutLogID string) error {
	query := `
		DELETE FROM workout_logs
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, workoutLogID)
	return err
}
