package user

import "github.com/google/uuid"

// User represents a user in the fitness and wellness application.
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Logged   bool   `json:"logged"`
}

// NewUser creates a new User instance.
func NewUser(email, password string) *User {
	return &User{
		ID:       uuid.New().String(),
		Email:    email,
		Password: password,
		Logged:   false,
	}
}
