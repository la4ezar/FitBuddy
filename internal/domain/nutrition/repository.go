package nutrition

import (
	"context"
	"database/sql"
	"github.com/FitBuddy/pkg/log"
	"time"
)

// Repository is a repository for managing nutrition log data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreateNutrition creates a new nutrition entry in the database.
func (r *Repository) CreateNutrition(ctx context.Context, nutrition *Nutrition) error {
	if _, err := r.db.ExecContext(ctx, `
        INSERT INTO nutrition (id, user_id, meal_id, grams, logged_at)
		VALUES ($1, (SELECT id FROM users WHERE email = $2), (SELECT id FROM meals WHERE name = $3), $4, $5)
    `, nutrition.ID, nutrition.UserEmail, nutrition.MealName, nutrition.Grams, nutrition.CreatedAt); err != nil {
		return err
	}

	return nil
}

// DeleteNutrition deletes a nutrition entry in the database.
func (r *Repository) DeleteNutrition(ctx context.Context, nutritionID string) error {
	if _, err := r.db.ExecContext(ctx, `
        DELETE FROM nutrition
		WHERE id = $1
    `, nutritionID); err != nil {
		return err
	}

	return nil
}

// GetAllNutritions retrieves all nutritions for user with email and date from the database.
func (r *Repository) GetAllNutritions(ctx context.Context, email string, date time.Time) ([]*Nutrition, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT n.id, u.email, m.name, n.grams, m.calories, n.logged_at
        FROM nutrition n
        JOIN users u ON n.user_id = u.id 
        JOIN meals m ON m.id = n.meal_id                               
		WHERE u.email = $1 AND DATE(n.logged_at) = DATE($2::timestamptz)
        ORDER BY n.logged_at DESC
    `, email, date)
	if err != nil {
		log.C(ctx).Info(err)
		return nil, err
	}
	defer rows.Close()

	var nutritions []*Nutrition

	for rows.Next() {
		var nutrition Nutrition
		if err := rows.Scan(&nutrition.ID, &nutrition.UserEmail, &nutrition.MealName, &nutrition.Grams, &nutrition.Calories, &nutrition.CreatedAt); err != nil {
			return nil, err
		}
		nutritions = append(nutritions, &nutrition)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nutritions, nil
}
