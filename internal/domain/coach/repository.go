package coach

import (
	"context"
	"database/sql"
	"fmt"
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

// GetAllCoaches fetches all coaches from the database.
func (r *Repository) GetAllCoaches(ctx context.Context) ([]*Coach, error) {
	query := "SELECT id, image_url, name, specialty FROM coaches"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var coaches []*Coach
	for rows.Next() {
		var coach Coach
		err := rows.Scan(&coach.ID, &coach.ImageURL, &coach.Name, &coach.Specialty)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		coaches = append(coaches, &coach)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return coaches, nil
}
