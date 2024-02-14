package forum

import (
	"context"
	"errors"
)

// Service handles business logic related to forum operations.
type Service struct {
	repository *Repository
}

// NewService creates a new Service instance.
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// CreatePost creates a new forum post.
func (s *Service) CreatePost(ctx context.Context, title, content, userEmail string) (*Post, error) {
	if title == "" || content == "" || userEmail == "" {
		return nil, errors.New("title, content, and user email are required")
	}

	newPost := NewPost(title, content, userEmail)

	if err := s.repository.CreatePost(ctx, newPost); err != nil {
		return nil, err
	}

	return newPost, nil
}

// GetAllPosts retrieves all forum posts.
func (s *Service) GetAllPosts(ctx context.Context) ([]*Post, error) {
	return s.repository.GetAllPosts(ctx)
}

// DeletePost deletes a forum post by ID.
func (s *Service) DeletePost(ctx context.Context, postID string) error {
	existingPost, err := s.repository.GetPostByID(ctx, postID)
	if err != nil {
		return err
	}
	if existingPost == nil {
		return errors.New("post not found")
	}

	return s.repository.DeletePost(ctx, postID)
}
