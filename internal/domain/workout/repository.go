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
        INSERT INTO workouts (id, user_id, exercise_id, sets, reps, weight, created_at)
		VALUES ($1, (SELECT id FROM users WHERE email = $2), (SELECT id FROM exercises WHERE name = $3), $4, $5, $6, $7)
    `, workout.ID, workout.UserEmail, workout.ExerciseName, workout.Sets, workout.Reps, workout.Weight, workout.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAllWorkouts(ctx context.Context, email string, date time.Time) ([]*Workout, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT w.id, u.email, e.name, w.sets, w.reps, w.weight, w.created_at
        FROM workouts w
        JOIN users u ON w.user_id = u.id 
        JOIN exercises e ON e.id = w.exercise_id                               
		WHERE u.email = $1 AND DATE(w.created_at) = DATE($2::timestamptz)
        ORDER BY w.created_at DESC
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

// GetLogByID retrieves a workout log entry from the database by ID.
func (r *Repository) GetLogByID(ctx context.Context, workoutLogID string) (*Workout, error) {
	query := `
		SELECT id, user_id, exercise, sets, reps, weight, logged_at
		FROM workout_logs
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, workoutLogID)

	var workoutLogItem Workout
	err := row.Scan(
		&workoutLogItem.ID,
		&workoutLogItem.UserEmail,
		&workoutLogItem.ExerciseName,
		&workoutLogItem.Sets,
		&workoutLogItem.Reps,
		&workoutLogItem.Weight,
		&workoutLogItem.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil // Workout log not found
	} else if err != nil {
		return nil, err
	}

	return &workoutLogItem, nil
}

// UpdateLog updates an existing workout log entry in the database.
func (r *Repository) UpdateLog(ctx context.Context, workoutLog *Workout) error {
	query := `
		UPDATE workout_logs
		SET exercise = $3, sets = $4, reps = $5, weight = $6, logged_at = $7
		WHERE id = $1 AND user_id = $2
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		workoutLog.ID,
		workoutLog.UserEmail,
		workoutLog.ExerciseName,
		workoutLog.Sets,
		workoutLog.Reps,
		workoutLog.Weight,
		workoutLog.CreatedAt,
	)

	return err
}

// DeleteLog deletes a workout log entry from the database by ID.
func (r *Repository) DeleteLog(ctx context.Context, workoutLogID string) error {
	query := `
		DELETE FROM workout_logs
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, workoutLogID)
	return err
}
