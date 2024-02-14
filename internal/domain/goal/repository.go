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
        INSERT INTO goals (id, user_id, name, description, start_date, end_date)
        VALUES ($1, (SELECT id FROM users WHERE email = $2), $3, $4, $5, $6)
        RETURNING id
    `
	_, err := r.db.ExecContext(ctx, query, goal.ID, email, goal.Name, goal.Description, goal.StartDate.Format(time.RFC3339), goal.EndDate.Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetGoalsByEmail(ctx context.Context, userEmail string) ([]*Goal, error) {
	query := "SELECT id, name, description, start_date, end_date FROM goals WHERE user_id = (SELECT id FROM users WHERE email = $1)"
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
		if err := rows.Scan(&goal.ID, &goal.Name, &goal.Description, &startDate, &endDate); err != nil {
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
		SELECT id, name, description, start_date, end_date
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
