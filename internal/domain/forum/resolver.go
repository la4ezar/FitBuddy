package forum

import (
	"context"
)

// Resolver handles GraphQL queries and mutations for the Forum aggregate.
type Resolver struct {
	forumService *Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(forumService *Service) *Resolver {
	return &Resolver{
		forumService: forumService,
	}
}

// CreatePostMutation is a GraphQL mutation to create a new forum post.
func (r *Resolver) CreatePostMutation(ctx context.Context, input Post) (*Post, error) {
	return r.forumService.CreatePost(ctx, input.Title, input.Content, input.AuthorID)
}

// CreateForumMutation is a GraphQL mutation to create a new forum.
func (r *Resolver) CreateForumMutation(ctx context.Context, input Forum) (*Forum, error) {
	return r.forumService.CreateForum(ctx, input.Name)
}

// GetPostQuery is a GraphQL query to retrieve a forum post by ID.
func (r *Resolver) GetPostQuery(ctx context.Context, postID string) (*Post, error) {
	return r.forumService.GetPostByID(ctx, postID)
}

// GetForumQuery is a GraphQL query to retrieve a forum by ID.
func (r *Resolver) GetForumQuery(ctx context.Context, forumID string) (*Forum, error) {
	return r.forumService.GetForumByID(ctx, forumID)
}

// UpdatePostMutation is a GraphQL mutation to update an existing forum post.
func (r *Resolver) UpdatePostMutation(ctx context.Context, input Post) (*Post, error) {
	return r.forumService.UpdatePost(ctx, input.ID, input.Title, input.Content)
}

// UpdateForumMutation is a GraphQL mutation to update an existing forum.
func (r *Resolver) UpdateForumMutation(ctx context.Context, input Forum) (*Forum, error) {
	return r.forumService.UpdateForum(ctx, input.ID, input.Name)
}

// DeletePostMutation is a GraphQL mutation to delete a forum post by ID.
func (r *Resolver) DeletePostMutation(ctx context.Context, postID string) (string, error) {
	err := r.forumService.DeletePost(ctx, postID)
	if err != nil {
		return "", err
	}
	return "Forum post deleted successfully", nil
}

// DeleteForumMutation is a GraphQL mutation to delete a forum by ID.
func (r *Resolver) DeleteForumMutation(ctx context.Context, forumID string) (string, error) {
	err := r.forumService.DeleteForum(ctx, forumID)
	if err != nil {
		return "", err
	}
	return "Forum deleted successfully", nil
}
