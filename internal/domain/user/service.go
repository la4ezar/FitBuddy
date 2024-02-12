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
func (s *Service) CreateUser(ctx context.Context, email, password string) (*User, error) {
	if email == "" || password == "" {
		return nil, errors.New("username and password are required")
	}

	existingUser, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email is already registered")
	}

	newUser := NewUser(email, password)

	if err := s.repository.CreateUser(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

// GetUserByEmail retrieves a user by ID.
func (s *Service) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return s.repository.GetUserByEmail(ctx, email)
}

// UpdateUser updates an existing user.
func (s *Service) UpdateUser(ctx context.Context, email, password string) (*User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}

	existingUser, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	existingUser.Email = email
	if password != "" {
		existingUser.Password = email
	}

	if err := s.repository.UpdateUser(ctx, existingUser); err != nil {
		return nil, err
	}

	return existingUser, nil
}

// LoginUser login an existing user.
func (s *Service) LoginUser(ctx context.Context, email, password string) (*User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}

	existingUser, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("user not found")
	}
	if existingUser.Password != password {
		return nil, errors.New("password don't match")
	}

	existingUser.Logged = true

	if err := s.repository.UpdateUser(ctx, existingUser); err != nil {
		return nil, err
	}

	return existingUser, nil
}

// LogoutUser logouts an existing user.
func (s *Service) LogoutUser(ctx context.Context, email string) (*User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}

	existingUser, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("user not found")
	}
	if !existingUser.Logged {
		return nil, errors.New("user is not logged in")
	}

	existingUser.Logged = false

	if err := s.repository.UpdateUser(ctx, existingUser); err != nil {
		return nil, err
	}

	return existingUser, nil
}

// DeleteUser deletes a user by ID.
func (s *Service) DeleteUser(ctx context.Context, email string) error {
	existingUser, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	return s.repository.DeleteUser(ctx, email)
}
