package sleep

import (
	"context"
	"database/sql"
	"time"
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
func (r *Repository) CreateLog(ctx context.Context, sleep *Log) error {
	startOfDay := time.Date(sleep.LoggedAt.Year(), sleep.LoggedAt.Month(), sleep.LoggedAt.Day(), 0, 0, 0, 0, sleep.LoggedAt.Location())

	_, err := r.db.ExecContext(ctx, "INSERT INTO sleep_logs (id, user_id, sleep_time, wake_time, logged_at) VALUES ($1, (SELECT id FROM users WHERE email = $2), $3, $4, $5)",
		sleep.ID, sleep.UserEmail, sleep.SleepTime.Format(time.RFC3339), sleep.WakeTime.Format(time.RFC3339), startOfDay.Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}

// GetSleepLogByEmailAndDate gets sleep by email and date
func (r *Repository) GetSleepLogByEmailAndDate(ctx context.Context, userEmail string, date time.Time) ([]*Log, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	rows, err := r.db.QueryContext(ctx, "SELECT id, sleep_time, wake_time, logged_at FROM sleep_logs WHERE user_id = (SELECT id FROM users WHERE email = $1) AND logged_at = $2::timestamp",
		userEmail, startOfDay.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sleeps []*Log

	for rows.Next() {
		var sleep Log
		err := rows.Scan(&sleep.ID, &sleep.SleepTime, &sleep.WakeTime, &sleep.LoggedAt)
		if err != nil {
			return nil, err
		}
		sleeps = append(sleeps, &sleep)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sleeps, nil
}

// GetLogByID retrieves a sleep log entry from the database by ID.
func (r *Repository) GetLogByID(ctx context.Context, sleepLogID string) (*Log, error) {
	query := `
		SELECT id, sleep_time, wake_time
		FROM sleep_logs
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, sleepLogID)

	var sleepLogItem Log
	err := row.Scan(
		&sleepLogItem.ID,
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
