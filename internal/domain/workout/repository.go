package workout

import (
	"context"
	"database/sql"
	"github.com/FitBuddy/pkg/log"
	"time"
)

// Repository is a repository for managing workout log data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreateWorkout creates a new workout entry in the database.
func (r *Repository) CreateWorkout(ctx context.Context, workout *Workout) error {
	if _, err := r.db.ExecContext(ctx, `
        INSERT INTO workouts (id, user_id, exercise_id, sets, reps, weight, logged_at)
		VALUES ($1, (SELECT id FROM users WHERE email = $2), (SELECT id FROM exercises WHERE name = $3), $4, $5, $6, $7)
    `, workout.ID, workout.UserEmail, workout.ExerciseName, workout.Sets, workout.Reps, workout.Weight, workout.CreatedAt); err != nil {
		return err
	}

	return nil
}

// DeleteWorkout deletes a workout entry in the database.
func (r *Repository) DeleteWorkout(ctx context.Context, workoutID string) error {
	if _, err := r.db.ExecContext(ctx, `
        DELETE FROM workouts
		WHERE id = $1
    `, workoutID); err != nil {
		return err
	}

	return nil
}

// GetAllWorkouts retrieves all workouts for user with email and date from the database.
func (r *Repository) GetAllWorkouts(ctx context.Context, email string, date time.Time) ([]*Workout, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT w.id, u.email, e.name, w.sets, w.reps, w.weight, w.logged_at
        FROM workouts w
        JOIN users u ON w.user_id = u.id 
        JOIN exercises e ON e.id = w.exercise_id                               
		WHERE u.email = $1 AND DATE(w.logged_at) = DATE($2::timestamptz)
        ORDER BY w.logged_at DESC
    `, email, date)
	if err != nil {
		log.C(ctx).Info(err)
		return nil, err
	}
	defer rows.Close()

	var workouts []*Workout

	for rows.Next() {
		var workout Workout
		if err := rows.Scan(&workout.ID, &workout.UserEmail, &workout.ExerciseName, &workout.Sets, &workout.Reps, &workout.Weight, &workout.CreatedAt); err != nil {
			return nil, err
		}
		workouts = append(workouts, &workout)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return workouts, nil
}
