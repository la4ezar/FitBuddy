package forum

import (
	"context"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
)

// Resolver handles GraphQL queries and mutations for the Forum aggregate.
type Resolver struct {
	service *Service
}

// NewResolver creates a new Resolver instance.
func NewResolver(service *Service) *Resolver {
	return &Resolver{
		service: service,
	}
}

// CreatePost is a GraphQL mutation to create a new forum post.
func (r *Resolver) CreatePost(ctx context.Context, title, content, email string) (*graphql.Post, error) {
	log.C(ctx).Infof("Creating post with title %q, content %q by user %q...", title, content, email)
	post, err := r.service.CreatePost(ctx, title, content, email)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Infof("Successfully created post with title %q, content %q by user %q", title, content, email)

	gqlPost := &graphql.Post{
		ID:        post.ID,
		UserEmail: post.UserEmail,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return gqlPost, nil
}

// DeletePost is a GraphQL mutation to delete a forum post.
func (r *Resolver) DeletePost(ctx context.Context, postID string) (bool, error) {
	log.C(ctx).Infof("Deleting post with ID %q...", postID)
	err := r.service.DeletePost(ctx, postID)
	if err != nil {
		return false, err
	}
	log.C(ctx).Infof("Successfully deleted post with ID %q", postID)

	return true, nil
}

// GetAllPosts is a GraphQL query to retrieve all forum posts.
func (r *Resolver) GetAllPosts(ctx context.Context) ([]*graphql.Post, error) {
	log.C(ctx).Info("Getting all posts...")
	posts, err := r.service.GetAllPosts(ctx)
	if err != nil {
		return nil, err
	}
	log.C(ctx).Info("Successfully got all posts")

	gqlPosts := make([]*graphql.Post, 0, len(posts))
	for _, p := range posts {
		gqlPosts = append(gqlPosts, &graphql.Post{
			ID:        p.ID,
			UserEmail: p.UserEmail,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return gqlPosts, nil
}
