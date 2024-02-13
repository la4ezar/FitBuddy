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

// CreateExercise creates a new exercise in the database.
func (r *Repository) CreateExercise(ctx context.Context, e *Exercise) error {
	query := `
		INSERT INTO exercises (id, name, description)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.ExecContext(ctx, query, e.ID, e.Name, e.Description)
	return err
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

// GetExerciseByID retrieves an exercise from the database by ID.
func (r *Repository) GetExerciseByID(ctx context.Context, exerciseID string) (*Exercise, error) {
	query := `
		SELECT id, name, description
		FROM exercises
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, exerciseID)

	var e Exercise
	err := row.Scan(&e.ID, &e.Name, &e.Description)
	if err == sql.ErrNoRows {
		return nil, nil // Exercise not found
	} else if err != nil {
		return nil, err
	}

	return &e, nil
}

// UpdateExercise updates an existing exercise in the database.
func (r *Repository) UpdateExercise(ctx context.Context, e *Exercise) error {
	query := `
		UPDATE exercises
		SET name = $2, description = $3
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, e.ID, e.Name, e.Description)
	return err
}

// DeleteExercise deletes an exercise from the database by ID.
func (r *Repository) DeleteExercise(ctx context.Context, exerciseID string) error {
	query := `
		DELETE FROM exercises
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, exerciseID)
	return err
}
