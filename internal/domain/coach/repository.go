package coach

import (
	"context"
	"database/sql"
)

// Repository is a repository for managing coach data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreateCoach creates a new coach in the database.
func (r *Repository) CreateCoach(ctx context.Context, coach *Coach) error {
	query := `
		INSERT INTO coaches (id, name, specialty)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.ExecContext(ctx, query, coach.ID, coach.Name, coach.Specialty)
	return err
}

// GetCoachByID retrieves a coach from the database by ID.
func (r *Repository) GetCoachByID(ctx context.Context, coachID string) (*Coach, error) {
	query := `
		SELECT id, name, specialty
		FROM coaches
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, coachID)

	var c Coach
	err := row.Scan(&c.ID, &c.Name, &c.Specialty)
	if err == sql.ErrNoRows {
		return nil, nil // Coach not found
	} else if err != nil {
		return nil, err
	}

	return &c, nil
}

// UpdateCoach updates an existing coach in the database.
func (r *Repository) UpdateCoach(ctx context.Context, coach *Coach) error {
	query := `
		UPDATE coaches
		SET name = $2, specialty = $3
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, coach.ID, coach.Name, coach.Specialty)
	return err
}

// DeleteCoach deletes a coach from the database by ID.
func (r *Repository) DeleteCoach(ctx context.Context, coachID string) error {
	query := `
		DELETE FROM coaches
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, coachID)
	return err
}
