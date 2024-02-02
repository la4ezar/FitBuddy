package forum

import "github.com/google/uuid"

// Forum represents an aggregation of forum posts in the application.
type Forum struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Posts []*Post `json:"posts"`
}

// NewForum creates a new Forum instance.
func NewForum(name string) *Forum {
	return &Forum{
		ID:    uuid.New().String(),
		Name:  name,
		Posts: make([]*Post, 0),
	}
}

// AddPost adds a new post to the forum.
func (f *Forum) AddPost(post *Post) {
	f.Posts = append(f.Posts, post)
}
