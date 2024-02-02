package user

import (
	"context"
	"errors"
)

// Resolver handles GraphQL queries and mutations for the User aggregate.
type Resolver struct {
	service *Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(service *Service) *Resolver {
	return &Resolver{
		service: service,
	}
}

// CreateUserMutation is a GraphQL mutation to create a new user.
func (r *Resolver) CreateUserMutation(ctx context.Context, input CreateUserInput) (*User, error) {
	return r.service.CreateUser(ctx, input.Username, input.Email)
}

// GetUserQuery is a GraphQL query to retrieve a user by ID.
func (r *Resolver) GetUserQuery(ctx context.Context, userID string) (*User, error) {
	return r.service.GetUserByID(ctx, userID)
}

// UpdateUserMutation is a GraphQL mutation to update an existing user.
func (r *Resolver) UpdateUserMutation(ctx context.Context, input UpdateUserInput) (*User, error) {
	return r.service.UpdateUser(ctx, input.UserID, input.Username, input.Email)
}

// DeleteUserMutation is a GraphQL mutation to delete a user by ID.
func (r *Resolver) DeleteUserMutation(ctx context.Context, userID string) (string, error) {
	if err := r.service.DeleteUser(ctx, userID); err != nil {
		return "", err
	}
	return "User deleted successfully", nil
}

// CreateUserInput represents the input for the CreateUserMutation.
type CreateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UpdateUserInput represents the input for the UpdateUserMutation.
type UpdateUserInput struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// GetUserByID retrieves a user by ID.
func (s *Service) GetUserByID(ctx context.Context, userID string) (*User, error) {
	return s.repository.GetUserByID(ctx, userID)
}

// UpdateUser updates an existing user.
func (s *Service) UpdateUser(ctx context.Context, userID, username, email string) (*User, error) {
	if username == "" || email == "" {
		return nil, errors.New("username and email are required")
	}

	existingUser, err := s.repository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	existingUser.Username = username
	existingUser.Email = email

	if err := s.repository.UpdateUser(ctx, existingUser); err != nil {
		return nil, err
	}

	return existingUser, nil
}

// DeleteUser deletes a user by ID.
func (s *Service) DeleteUser(ctx context.Context, userID string) error {
	existingUser, err := s.repository.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	return s.repository.DeleteUser(ctx, userID)
}
