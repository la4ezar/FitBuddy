package exercise

import (
	"context"
	"database/sql"
)

// Repository is a repository for managing exercise data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// GetAllExercises retrieves all exercises from the database.
func (r *Repository) GetAllExercises(ctx context.Context) ([]*Exercise, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT e.id, e.name FROM exercises e
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []*Exercise

	for rows.Next() {
		var exercise Exercise
		if err := rows.Scan(&exercise.ID, &exercise.Name); err != nil {
			return nil, err
		}
		exercises = append(exercises, &exercise)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exercises, nil
}
