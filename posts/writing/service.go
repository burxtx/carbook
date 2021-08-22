package writing

import (
	"context"

	"github.com/burxtx/car/posts"
)

// Service is the interface that provides posts related methods.
type Service interface {
	Create(ctx context.Context, author, vehicle, title, excerpt, body string) (*posts.Post, error)
}

type service struct {
	repo posts.PostRepository
}

func New(repo posts.PostRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, author, vehicle, title, excerpt, body string) (*posts.Post, error) {
	p := posts.NewPost(author, vehicle, title, excerpt, body)
	err := s.repo.Create(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Post is a real model for view
type Post struct {
	PostID  string
	Title   string
	Excerpt string
	Body    string
}
