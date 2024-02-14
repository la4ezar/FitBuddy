package forum

import (
	"context"
	"database/sql"
)

// Repository is a repository for managing forum data.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreatePost creates a new forum post in the database.
func (r *Repository) CreatePost(ctx context.Context, post *Post) error {
	if _, err := r.db.ExecContext(ctx, `
        INSERT INTO posts (id, user_id, title, content, logged_at)
        VALUES ($1, (SELECT id FROM users WHERE email = $2), $3, $4, $5)
    `, post.ID, post.UserEmail, post.Title, post.Content, post.CreatedAt); err != nil {
		return err
	}

	return nil
}

// GetPostByID retrieves a forum post from the database by ID.
func (r *Repository) GetPostByID(ctx context.Context, postID string) (*Post, error) {
	query := `
		SELECT id, title, content
		FROM posts
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, postID)

	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Content)
	if err == sql.ErrNoRows {
		return nil, nil // Post not found
	} else if err != nil {
		return nil, err
	}

	return &post, nil
}

// GetAllPosts retrieves all posts from the database.
func (r *Repository) GetAllPosts(ctx context.Context) ([]*Post, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT p.id, u.email, p.title, p.content, p.logged_at
        FROM posts p
        JOIN users u ON p.user_id = u.id
        ORDER BY p.logged_at DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post

	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.UserEmail, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// DeletePost deletes a forum post from the database by ID.
func (r *Repository) DeletePost(ctx context.Context, postID string) error {
	query := `
		DELETE FROM posts
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, postID)
	return err
}
