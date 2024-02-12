package user

import (
	"context"
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
	return r.service.CreateUser(ctx, input.Email, input.Password)
}

// GetUserQuery is a GraphQL query to retrieve a user by ID.
func (r *Resolver) GetUserQuery(ctx context.Context, email string) (*User, error) {
	return r.service.GetUserByEmail(ctx, email)
}

// UpdateUserMutation is a GraphQL mutation to update an existing user.
func (r *Resolver) UpdateUserMutation(ctx context.Context, input UpdateUserInput) (*User, error) {
	return r.service.UpdateUser(ctx, input.Email, input.Password)
}

func (r *Resolver) LoginUserMutation(ctx context.Context, email, password string) (*User, error) {
	return r.service.LoginUser(ctx, email, password)
}

func (r *Resolver) LogoutUserMutation(ctx context.Context, email string) (*User, error) {
	return r.service.LogoutUser(ctx, email)
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
	Email    string `json:"email"`
	Password string `json:"password"`
	Logged   bool   `json:"logged"`
}

// UpdateUserInput represents the input for the UpdateUserMutation.
type UpdateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Logged   bool   `json:"logged"`
}
