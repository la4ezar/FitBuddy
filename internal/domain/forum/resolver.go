package forum

import (
	"context"
	"github.com/FitBuddy/pkg/graphql"
	"github.com/FitBuddy/pkg/log"
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

// CreatePost is a GraphQL mutation to create a new forum post.
func (r *Resolver) CreatePost(ctx context.Context, title, content, email string) (*graphql.Post, error) {
	log.C(ctx).Infof("Creating post with title %q, content %q by user %q...", title, content, email)
	post, err := r.forumService.CreatePost(ctx, title, content, email)
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

// CreateForumMutation is a GraphQL mutation to create a new forum.
func (r *Resolver) CreateForumMutation(ctx context.Context, input Forum) (*Forum, error) {
	return r.forumService.CreateForum(ctx, input.Name)
}

// GetAllPosts is a GraphQL query to retrieve all forum posts.
func (r *Resolver) GetAllPosts(ctx context.Context) ([]*graphql.Post, error) {
	log.C(ctx).Info("Getting all posts...")
	posts, err := r.forumService.GetAllPosts(ctx)
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
