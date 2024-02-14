package booking

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Repository is a repository for managing bookings data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
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
