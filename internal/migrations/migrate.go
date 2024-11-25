package migrations

import (
	"context"
	"time"

	"github.com/VandiKond/vanerrors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Users struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type Video struct {
	ID          int64     `json:"id"`
	AuthorID    int64     `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Views       int64     `json:"views"`
	Likes       int64     `json:"likes"`
	Thumbnail   string    `json:"thumbnail"`
	Filepath    string    `json:"filepath"`
}

type Comments struct {
	ID        int64     `json:"id"`
	AuthorID  int64     `json:"author"`
	VideoID   int64     `json:"video_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Likes     int64     `json:"likes"`
}

func InitTables(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), `
     CREATE TABLE IF NOT EXISTS users (
          id SERIAL PRIMARY KEY,
          username VARCHAR(255) UNIQUE NOT NULL,
          email VARCHAR(255) UNIQUE NOT NULL,
          password VARCHAR(255) NOT NULL,
          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
     );

     CREATE TABLE IF NOT EXISTS videos (
          id SERIAL PRIMARY KEY,
          author_id INTEGER NOT NULL,
          title VARCHAR(255) NOT NULL,
          description TEXT,
          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
          views INTEGER NOT NULL DEFAULT 0,
          thumbnail VARCHAR(255) NOT NULL,
          filepath VARCHAR(255) NOT NULL
     );

     CREATE TABLE IF NOT EXISTS comments (
          id SERIAL PRIMARY KEY,
		author_id INTEGER REFERENCES users(id) NOT NULL,
          video_id INTEGER REFERENCES videos(id) NOT NULL,
          content TEXT NOT NULL,
          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
          likes INTEGER NOT NULL DEFAULT 0
	);
	CREATE TABLE IF NOT EXISTS likes (
	     id SERIAL PRIMARY KEY,
          user_id INTEGER REFERENCES users(id) NOT NULL,
          video_id INTEGER REFERENCES videos(id) NOT NULL,
          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP);
	`)
	if err != nil {
		err = vanerrors.NewWrap("unable to create tables", err, vanerrors.EmptyHandler)
		return err
	}
	return nil
}
