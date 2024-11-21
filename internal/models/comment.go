package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	Author    User      `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time    `json:"created_at"`
	Likes     int       `json:"likes"`
	Replies   []Comment `json:"replies"`
}