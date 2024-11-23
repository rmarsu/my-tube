package models

import "time"

type Video struct {
	ID          int     `json:"id"`
	AuthorID    int     `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Views       int     `json:"views"`
	Likes       int     `json:"likes"`
	Thumbnail   string    `json:"thumbnail"`
	Filepath    string    `json:"filepath"`
}