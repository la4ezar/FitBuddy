package coach

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
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

// BookCoach checks if a coach is booked in the database.
func (r *Repository) BookCoach(ctx context.Context, email, coachName string) (bool, error) {
	existingBookingQuery := "SELECT 1 FROM bookings WHERE user_id = (SELECT id FROM users WHERE email = $1) AND coach_id = (SELECT id FROM coaches WHERE name = $2)"
	var count int
	err := r.db.QueryRowContext(ctx, existingBookingQuery, email, coachName).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if count > 0 {
		return false, errors.New("user is already booked with the coach")
	}

	// Insert a new booking record
	insertBookingQuery := "INSERT INTO bookings (id, user_id, coach_id) VALUES ($1, (SELECT id FROM users WHERE email = $2), (SELECT id FROM coaches WHERE name = $3))"
	if _, err := r.db.ExecContext(ctx, insertBookingQuery, uuid.New().String(), email, coachName); err != nil {
		return false, err
	}

	return true, nil
}

// UnbookCoach unbooks a coach for a user
func (r *Repository) UnbookCoach(ctx context.Context, email, name string) (bool, error) {
	deleteBookingQuery := "DELETE FROM bookings WHERE user_id = (SELECT id FROM users WHERE email = $1) AND coach_id = (SELECT id FROM coaches WHERE name = $2)"
	if _, err := r.db.ExecContext(ctx, deleteBookingQuery, email, name); err != nil {
		return false, err
	}
	return true, nil
}

// IsCoachBookedByUser checks if coach with given name is booked by a user with given email
func (r *Repository) IsCoachBookedByUser(ctx context.Context, coachName, email string) (bool, error) {
	existingBookingQuery := "SELECT 1 FROM bookings WHERE user_id = (SELECT id FROM users WHERE email = $1) AND coach_id = (SELECT id FROM coaches WHERE name = $2)"
	var count int
	err := r.db.QueryRowContext(ctx, existingBookingQuery, email, coachName).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return count > 0, nil
}

// IsCoachBooked checks if coach with given name is booked
func (r *Repository) IsCoachBooked(ctx context.Context, coachName string) (bool, error) {
	existingBookingQuery := "SELECT 1 FROM bookings WHERE coach_id = (SELECT id FROM coaches WHERE name = $1)"
	var count int
	err := r.db.QueryRowContext(ctx, existingBookingQuery, coachName).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return count > 0, nil
}
