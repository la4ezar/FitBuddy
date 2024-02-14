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

// CreateForum creates a new forum in the database.
func (r *Repository) CreateForum(ctx context.Context, forum *Forum) error {
	query := `
		INSERT INTO forums (id, name)
		VALUES ($1, $2)
	`

	_, err := r.db.ExecContext(ctx, query, forum.ID, forum.Name)
	return err
}

// GetPostByID retrieves a forum post from the database by ID.
func (r *Repository) GetPostByID(ctx context.Context, postID string) (*Post, error) {
	query := `
		SELECT id, title, content, author_id, logged_at
		FROM forum_posts
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, postID)

	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.UserEmail, &post.CreatedAt)
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

// GetForumByID retrieves a forum from the database by ID.
func (r *Repository) GetForumByID(ctx context.Context, forumID string) (*Forum, error) {
	query := `
		SELECT id, name
		FROM forums
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, forumID)

	var forum Forum
	err := row.Scan(&forum.ID, &forum.Name)
	if err == sql.ErrNoRows {
		return nil, nil // Forum not found
	} else if err != nil {
		return nil, err
	}

	return &forum, nil
}

// UpdatePost updates an existing forum post in the database.
func (r *Repository) UpdatePost(ctx context.Context, post *Post) error {
	query := `
		UPDATE forum_posts
		SET title = $2, content = $3
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, post.ID, post.Title, post.Content)
	return err
}

// UpdateForum updates an existing forum in the database.
func (r *Repository) UpdateForum(ctx context.Context, forum *Forum) error {
	query := `
		UPDATE forums
		SET name = $2
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, forum.ID, forum.Name)
	return err
}

// DeletePost deletes a forum post from the database by ID.
func (r *Repository) DeletePost(ctx context.Context, postID string) error {
	query := `
		DELETE FROM forum_posts
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, postID)
	return err
}

// DeleteForum deletes a forum from the database by ID.
func (r *Repository) DeleteForum(ctx context.Context, forumID string) error {
	query := `
		DELETE FROM forums
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, forumID)
	return err
}
