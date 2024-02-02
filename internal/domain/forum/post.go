package forum

import (
	"github.com/google/uuid"
	"time"
)

// Post represents an individual forum post in the application.
type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  string    `json:"authorId"`
	CreatedAt time.Time `json:"createdAt"`
}

// NewPost creates a new Post instance.
func NewPost(title, content, authorID string) *Post {
	return &Post{
		ID:        uuid.New().String(),
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		CreatedAt: time.Now(),
	}
}
