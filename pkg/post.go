package pkg

import (
	"errors"
	"time"
)

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"name"`
	Body      string `json:"body"`
	CreatedAt int64  `json:"created_at"`
}

type PostDTO struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewPost(id int, title, body string) (*Post, error) {
	if title == "" || body == "" {
		return nil, errors.New("the fields 'title' and 'body' are required")
	}

	if len(title) < 5 {
		return nil, errors.New("the 'title' field is too short. Minimum is 5 characters")
	}

	if len(body) < 10 {
		return nil, errors.New("the 'body' field is too short. Minimum is 10 characters")
	}

	post := new(Post)

	post.ID = id
	post.Title = title
	post.Body = body
	post.CreatedAt = time.Now().UnixNano()

	return post, nil
}
