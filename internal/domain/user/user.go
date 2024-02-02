package user

import "github.com/google/uuid"

// User represents a user in the fitness and wellness application.
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// NewUser creates a new User instance.
func NewUser(username, email string) *User {
	return &User{
		ID:       generateUniqueID(),
		Username: username,
		Email:    email,
	}
}

// generateUniqueID generates UUID
func generateUniqueID() string {
	return uuid.New().String()
}
