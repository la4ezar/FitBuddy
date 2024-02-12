package user

import (
	"context"
	"database/sql"
)

// Repository is a repository for managing user data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new UserRepository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreateUser creates a new user in the database.
func (r *Repository) CreateUser(ctx context.Context, u *User) error {
	query := `
		INSERT INTO users (id, email, password, logged)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(ctx, query, u.ID, u.Email, u.Password, u.Logged)
	return err
}

//// GetUserByID retrieves a user from the database by ID.
//func (r *Repository) GetUserByID(ctx context.Context, userID string) (*User, error) {
//	query := `
//		SELECT id, email, password, logged
//		FROM users
//		WHERE id = $1
//	`
//
//	row := r.db.QueryRowContext(ctx, query, userID)
//
//	var u User
//	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Logged)
//	if err == sql.ErrNoRows {
//		return nil, nil // User not found
//	} else if err != nil {
//		return nil, err
//	}
//
//	return &u, nil
//}

// GetUserByEmail retrieves a user from the database by email.
func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, email, password, logged
		FROM users
		WHERE email = $1
	`

	row := r.db.QueryRowContext(ctx, query, email)

	var u User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Logged)
	if err == sql.ErrNoRows {
		return nil, nil // User not found
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}

// UpdateUser updates an existing user in the database.
func (r *Repository) UpdateUser(ctx context.Context, u *User) error {
	query := `
		UPDATE users
		SET email = $2, password = $3, logged = $4
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, u.ID, u.Email, u.Password, u.Logged)
	return err
}

// DeleteUser deletes a user from the database by email.
func (r *Repository) DeleteUser(ctx context.Context, email string) error {
	query := `
		DELETE FROM users
		WHERE email = $1
	`

	_, err := r.db.ExecContext(ctx, query, email)
	return err
}
