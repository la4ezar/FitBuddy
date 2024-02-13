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

// CreateForum creates a new forum.
func (s *Service) CreateForum(ctx context.Context, name string) (*Forum, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	newForum := NewForum(name)

	if err := s.repository.CreateForum(ctx, newForum); err != nil {
		return nil, err
	}

	return newForum, nil
}

// GetPostByID retrieves a forum post by ID.
func (s *Service) GetPostByID(ctx context.Context, postID string) (*Post, error) {
	return s.repository.GetPostByID(ctx, postID)
}

// GetAllPosts retrieves all forum posts.
func (s *Service) GetAllPosts(ctx context.Context) ([]*Post, error) {
	return s.repository.GetAllPosts(ctx)
}

// GetForumByID retrieves a forum by ID.
func (s *Service) GetForumByID(ctx context.Context, forumID string) (*Forum, error) {
	return s.repository.GetForumByID(ctx, forumID)
}

// UpdatePost updates an existing forum post.
func (s *Service) UpdatePost(ctx context.Context, postID, title, content string) (*Post, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content are required")
	}

	existingPost, err := s.repository.GetPostByID(ctx, postID)
	if err != nil {
		return nil, err
	}
	if existingPost == nil {
		return nil, errors.New("post not found")
	}

	existingPost.Title = title
	existingPost.Content = content

	if err := s.repository.UpdatePost(ctx, existingPost); err != nil {
		return nil, err
	}

	return existingPost, nil
}

// UpdateForum updates an existing forum.
func (s *Service) UpdateForum(ctx context.Context, forumID, name string) (*Forum, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	existingForum, err := s.repository.GetForumByID(ctx, forumID)
	if err != nil {
		return nil, err
	}
	if existingForum == nil {
		return nil, errors.New("forum not found")
	}

	existingForum.Name = name

	if err := s.repository.UpdateForum(ctx, existingForum); err != nil {
		return nil, err
	}

	return existingForum, nil
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

// DeleteForum deletes a forum by ID.
func (s *Service) DeleteForum(ctx context.Context, forumID string) error {
	existingForum, err := s.repository.GetForumByID(ctx, forumID)
	if err != nil {
		return err
	}
	if existingForum == nil {
		return errors.New("forum not found")
	}

	return s.repository.DeleteForum(ctx, forumID)
}
