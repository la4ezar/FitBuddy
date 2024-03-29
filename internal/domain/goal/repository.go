package goal

import (
	"context"
	"database/sql"
	"time"
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

// CreateGoal creates a new goal for a user with the given email
func (r *Repository) CreateGoal(ctx context.Context, email string, goal *Goal) error {
	query := `
        INSERT INTO goals (id, user_id, name, description, start_date, end_date, completed)
        VALUES ($1, (SELECT id FROM users WHERE email = $2), $3, $4, $5, $6, $7)
        RETURNING id
    `
	_, err := r.db.ExecContext(ctx, query, goal.ID, email, goal.Name, goal.Description, goal.StartDate.Format(time.RFC3339), goal.EndDate.Format(time.RFC3339), goal.Completed)
	if err != nil {
		return err
	}

	return nil
}

// GetGoalsByEmail gets all goals by user email from the database.
func (r *Repository) GetGoalsByEmail(ctx context.Context, userEmail string) ([]*Goal, error) {
	query := "SELECT id, name, description, start_date, end_date, completed FROM goals WHERE user_id = (SELECT id FROM users WHERE email = $1)"
	rows, err := r.db.QueryContext(ctx, query, userEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []*Goal
	var startDate string
	var endDate string
	for rows.Next() {
		var goal Goal
		if err := rows.Scan(&goal.ID, &goal.Name, &goal.Description, &startDate, &endDate, &goal.Completed); err != nil {
			return nil, err
		}

		goal.StartDate, err = time.Parse(time.RFC3339, startDate)
		if err != nil {
			return nil, err
		}
		goal.EndDate, err = time.Parse(time.RFC3339, endDate)
		if err != nil {
			return nil, err
		}

		goals = append(goals, &goal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return goals, nil
}

// GetGoalByID retrieves a fitness or wellness goal from the database by ID.
func (r *Repository) GetGoalByID(ctx context.Context, goalID string) (*Goal, error) {
	query := `
		SELECT id, name, description, start_date, end_date, completed
		FROM goals
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, goalID)

	var goalItem Goal
	err := row.Scan(
		&goalItem.ID,
		&goalItem.Name,
		&goalItem.Description,
		&goalItem.StartDate,
		&goalItem.EndDate,
		&goalItem.Completed,
	)
	if err == sql.ErrNoRows {
		return nil, nil // Goal not found
	} else if err != nil {
		return nil, err
	}

	return &goalItem, nil
}

// CompleteGoalByID updates goal completed field to true.
func (r *Repository) CompleteGoalByID(ctx context.Context, goalID string) error {
	query := `
		UPDATE goals
		SET completed = true
		WHERE id = $1
	`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		goalID,
	); err != nil {
		return err
	}

	return nil
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
