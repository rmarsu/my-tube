package models

import "time"

type Comment struct {
	ID        int64     `json:"id"`
	AuthorID  int64     `json:"author"`
	VideoID   int64     `json:"video_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Likes     int64     `json:"likes"`
}