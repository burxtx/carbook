package posts

import (
	"github.com/burxtx/car/posts/utils"
)

type PostID string

// Post is the central of post domain model
// play as aggragate root
type Post struct {
	ID      PostID
	Author  string
	Vehicle string
	Cover   string
	Title   string
	Excerpt string
	Body    string
}

// NewUser creates a new user
func NewPost(author, vehicle, title, excerpt, body string) *Post {
	return &Post{
		ID:      PostID(utils.NewPostID()),
		Author:  author,
		Vehicle: vehicle,
		Title:   title,
		Excerpt: excerpt,
		Body:    body,
	}
}

// Repository provides access to a store.
type PostRepository interface {
	Create(post *Post) error
	Find(postID string) (*Post, error)
	List() []*Post
}
