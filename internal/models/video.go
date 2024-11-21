package models

import "time"

type Video struct {
	ID          int       `json:"id"`
	Author      User      `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Views       int       `json:"views"`
	Likes       int       `json:"likes"`
	Comments    []Comment `json:"comments"`

	Thumbnail string `json:"thumbnail"`
	Filepath  string `json:"filepath"`
}
