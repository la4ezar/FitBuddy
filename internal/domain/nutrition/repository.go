package nutrition

import (
	"context"
	"database/sql"
)

// Repository is a repository for managing nutrition log data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new NutritionRepository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreateLog creates a new nutrition log entry in the database.
func (r *Repository) CreateLog(ctx context.Context, nutritionLog *Log) error {
	query := `
		INSERT INTO nutrition_logs (id, user_id, meal, description, calories, logged_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		nutritionLog.ID,
		nutritionLog.UserID,
		nutritionLog.Meal,
		nutritionLog.Description,
		nutritionLog.Calories,
		nutritionLog.LoggedAt,
	)

	return err
}

// GetLogByID retrieves a nutrition log entry from the database by ID.
func (r *Repository) GetLogByID(ctx context.Context, nutritionLogID string) (*Log, error) {
	query := `
		SELECT id, user_id, meal, description, calories, logged_at
		FROM nutrition_logs
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, nutritionLogID)

	var nutritionLogItem Log
	err := row.Scan(
		&nutritionLogItem.ID,
		&nutritionLogItem.UserID,
		&nutritionLogItem.Meal,
		&nutritionLogItem.Description,
		&nutritionLogItem.Calories,
		&nutritionLogItem.LoggedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil // Nutrition log not found
	} else if err != nil {
		return nil, err
	}

	return &nutritionLogItem, nil
}

// UpdateLog updates an existing nutrition log entry in the database.
func (r *Repository) UpdateLog(ctx context.Context, nutritionLog *Log) error {
	query := `
		UPDATE nutrition_logs
		SET meal = $2, description = $3, calories = $4, logged_at = $5
		WHERE id = $1
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		nutritionLog.ID,
		nutritionLog.Meal,
		nutritionLog.Description,
		nutritionLog.Calories,
		nutritionLog.LoggedAt,
	)

	return err
}

// DeleteLog deletes a nutrition log entry from the database by ID.
func (r *Repository) DeleteLog(ctx context.Context, nutritionLogID string) error {
	query := `
		DELETE FROM nutrition_logs
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, nutritionLogID)
	return err
}
