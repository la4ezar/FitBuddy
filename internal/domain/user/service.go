package user

import (
	"context"
	"errors"
)

// Service handles business logic related to user operations.
type Service struct {
	repository *Repository
}

// NewService creates a new Service instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// CreateUser creates a new user.
func (s *Service) CreateUser(ctx context.Context, username, email string) (*User, error) {
	if username == "" || email == "" {
		return nil, errors.New("username and email are required")
	}

	existingUser, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email is already registered")
	}

	newUser := NewUser(username, email)

	if err := s.repository.CreateUser(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}
