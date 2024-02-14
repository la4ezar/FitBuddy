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

// NewRepository creates a new NutritionRepository instance.
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

// GetAllMeals retrieves all exercises from the database.
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
