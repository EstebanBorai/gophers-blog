package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

type Repository struct {
	posts    map[int]Post
	randSeed rand.Rand
}

type DataSource interface {
	DeletePost(id int) (*Post, error)
	FindPost(id int) (*Post, error)
	InsertPost(dto *PostDTO) (*Post, error)
	ShowPosts() ([]Post, error)
	UpdatePost(id int, dto *PostDTO) (*Post, error)
}

func NewRepository() Repository {
	randSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	posts := make(map[int]Post)
	repos := Repository{
		randSeed: *randSeed,
		posts:    posts,
	}

	return repos
}

func (r *Repository) newId() int {
	return r.randSeed.Int()
}

func (r *Repository) ShowPosts() ([]Post, error) {
	if len(r.posts) == 0 {
		return make([]Post, 0), nil
	}

	list := make([]Post, len(r.posts))
	idx := 0

	for _, entry := range r.posts {
		list[idx] = entry
		idx++
	}

	return list, nil
}

func (r *Repository) FindPost(id int) (*Post, error) {
	if post, found := r.posts[id]; found {
		return &post, nil
	}

	return nil, fmt.Errorf("no post with id: %d found", id)
}

func (r *Repository) InsertPost(dto *PostDTO) (*Post, error) {
	id := r.newId()

	if post, err := NewPost(id, dto.Title, dto.Body); err != nil {
		return nil, err
	} else {
		r.posts[id] = *post

		return post, nil
	}
}

func (r *Repository) UpdatePost(id int, dto *PostDTO) (*Post, error) {
	if post, err := r.FindPost(id); err != nil {
		return nil, err
	} else {
		// TODO: Create a function to validation Post fields
		if tmp, err := NewPost(id, post.Title, post.Body); err != nil {
			return nil, err
		} else {
			tmp.ID = id
			tmp.CreatedAt = post.CreatedAt

			r.posts[id] = *tmp

			return tmp, nil
		}
	}
}

func (r *Repository) DeletePost(id int) (*Post, error) {
	if post, err := r.FindPost(id); err != nil {
		return nil, err
	} else {
		delete(r.posts, id)

		return post, nil
	}
}
