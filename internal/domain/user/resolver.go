package user

import (
	"context"
	"github.com/FitBuddy/internal/domain/leaderboard"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
)

// Resolver handles GraphQL queries and mutations for the User aggregate.
type Resolver struct {
	service            *Service
	leaderBoardService *leaderboard.Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(service *Service, leaderboardService *leaderboard.Service) *Resolver {
	return &Resolver{
		service:            service,
		leaderBoardService: leaderboardService,
	}
}

// CreateUser is a GraphQL mutation to create a new user.
func (r *Resolver) CreateUser(ctx context.Context, email, password string) (*graphql.User, error) {
	log.C(ctx).Infof("Creating User with email %q...", email)

	u, err := r.service.CreateUser(ctx, email, password)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully created user with email %q", email)

	log.C(ctx).Infof("Creating record in leaderboard for user with email %q...", email)

	err = r.leaderBoardService.Create(ctx, email, 0)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully created record in leaderboard for user with email %q", email)

	gqlUser := &graphql.User{
		ID:    u.ID,
		Email: u.Email,
	}
	return gqlUser, nil
}

// GetUserQuery is a GraphQL query to retrieve a user by ID.
func (r *Resolver) GetUserQuery(ctx context.Context, email string) (*User, error) {
	return r.service.GetUserByEmail(ctx, email)
}

// UpdateUserMutation is a GraphQL mutation to update an existing user.
func (r *Resolver) UpdateUserMutation(ctx context.Context, input UpdateUserInput) (*User, error) {
	return r.service.UpdateUser(ctx, input.Email, input.Password)
}

func (r *Resolver) LoginUser(ctx context.Context, email, password string) (*graphql.User, error) {
	log.C(ctx).Infof("Logging user with email %q...", email)
	u, err := r.service.LoginUser(ctx, email, password)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully logged user with email %q...", email)

	gqlUser := &graphql.User{
		ID:    u.ID,
		Email: u.Email,
	}
	return gqlUser, nil
}

func (r *Resolver) LogoutUser(ctx context.Context, email string) (*graphql.User, error) {
	log.C(ctx).Infof("Logging out user with email %q...", email)
	u, err := r.service.LogoutUser(ctx, email)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully logging out user with email %q...", email)

	gqlUser := &graphql.User{
		ID:    u.ID,
		Email: u.Email,
	}
	return gqlUser, nil
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
