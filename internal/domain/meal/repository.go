package meal

import (
	"context"
	"database/sql"
)

// Repository is a repository for managing nutrition meal data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Meal Repository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// GetAllMeals retrieves all meals from the database.
func (r *Repository) GetAllMeals(ctx context.Context) ([]*Meal, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT m.id, m.name FROM meals m
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meals []*Meal

	for rows.Next() {
		var meal Meal
		if err := rows.Scan(&meal.ID, &meal.Name); err != nil {
			return nil, err
		}
		meals = append(meals, &meal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return meals, nil
}
