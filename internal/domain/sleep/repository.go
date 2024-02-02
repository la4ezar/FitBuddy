package sleep

import (
	"context"
	"database/sql"
)

// Repository is a repository for managing sleep log data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new SleepRepository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreateLog creates a new sleep log entry in the database.
func (r *Repository) CreateLog(ctx context.Context, sleepLog *Log) error {
	query := `
		INSERT INTO sleep_logs (id, user_id, duration, sleep_time, wake_time)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		sleepLog.ID,
		sleepLog.UserID,
		sleepLog.Duration,
		sleepLog.SleepTime,
		sleepLog.WakeTime,
	)

	return err
}

// GetLogByID retrieves a sleep log entry from the database by ID.
func (r *Repository) GetLogByID(ctx context.Context, sleepLogID string) (*Log, error) {
	query := `
		SELECT id, user_id, duration, sleep_time, wake_time
		FROM sleep_logs
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, sleepLogID)

	var sleepLogItem Log
	err := row.Scan(
		&sleepLogItem.ID,
		&sleepLogItem.UserID,
		&sleepLogItem.Duration,
		&sleepLogItem.SleepTime,
		&sleepLogItem.WakeTime,
	)
	if err == sql.ErrNoRows {
		return nil, nil // Sleep log not found
	} else if err != nil {
		return nil, err
	}

	return &sleepLogItem, nil
}

// UpdateLog updates an existing sleep log entry in the database.
func (r *Repository) UpdateLog(ctx context.Context, sleepLog *Log) error {
	query := `
		UPDATE sleep_logs
		SET duration = $2, sleep_time = $3, wake_time = $4
		WHERE id = $1
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		sleepLog.ID,
		sleepLog.Duration,
		sleepLog.SleepTime,
		sleepLog.WakeTime,
	)

	return err
}

// DeleteLog deletes a sleep log entry from the database by ID.
func (r *Repository) DeleteLog(ctx context.Context, sleepLogID string) error {
	query := `
		DELETE FROM sleep_logs
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, sleepLogID)
	return err
}
