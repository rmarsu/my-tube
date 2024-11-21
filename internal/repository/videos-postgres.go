package repository

import (
	"context"
	"fmt"
	"myTube/internal/models"

	"github.com/jackc/pgx/v5"
)

type VideosRepo struct {
	db *pgx.Conn
}

func NewVideoRepository(db *pgx.Conn) *VideosRepo {
	return &VideosRepo{db: db}
}

func (r *VideosRepo) GetById(ctx context.Context, id int) (models.Video, error) {
	var video models.Video
	query := `SELECT id, author, title, description, created_at, views, likes, comments, thumbnail, filepath FROM videos WHERE id = $1`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&video.ID,
		&video.Author.Username,
		&video.Title,
		&video.Description,
		&video.CreatedAt,
		&video.Views,
		&video.Likes,
		&video.Comments,
		&video.Thumbnail,
		&video.Filepath,
	)
	if err != nil {
		return models.Video{}, err
	}

	// Increment the view count
	updateQuery := `UPDATE videos SET views = views + 1 WHERE id = $1`
	_, err = r.db.Exec(ctx, updateQuery, id)
	if err != nil {
		return video, fmt.Errorf("error updating view count: %v", err)
	} else {
		video.Views++
	}

	return video, nil
}

func (r *VideosRepo) GetTrendy(ctx context.Context) ([]models.Video, error) {
	var videos []models.Video
	query := `SELECT id, author, title, description, created_at, views, likes, comments, thumbnail, filepath FROM videos ORDER BY likes DESC LIMIT 10`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var video models.Video
		err := rows.Scan(
			&video.ID,
			&video.Author.Username,
			&video.Title,
			&video.Description,
			&video.CreatedAt,
			&video.Views,
			&video.Likes,
			&video.Comments,
			&video.Thumbnail,
			&video.Filepath,
		)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func (r *VideosRepo) Create(ctx context.Context, video models.Video) error {
	query := `INSERT INTO videos (author, title, description, created_at, views, likes, comments, thumbnail, filepath) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err := r.db.QueryRow(ctx, query, video.Author.Username, video.Title, video.Description, video.CreatedAt, video.Views, video.Likes, video.Comments, video.Thumbnail, video.Filepath).Scan(&video.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *VideosRepo) Update(ctx context.Context, id int, video models.Video) error {
	query := `UPDATE videos SET title=$1, description=$2, thumbnail=$3, WHERE id=$4`

	_, err := r.db.Exec(ctx, query, video.Title, video.Description, video.Thumbnail, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *VideosRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM videos WHERE id=$1`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *VideosRepo) Search(ctx context.Context, query string) ([]models.Video, error) {
	var videos []models.Video
	searchQuery := fmt.Sprintf(`SELECT id, author, title, description, created_at, views, likes, comments, thumbnail, filepath FROM videos WHERE LOWER(title) LIKE LOWER('%%%s%%') OR LOWER(description) LIKE LOWER('%%%s%%')`, query, query)

	rows, err := r.db.Query(ctx, searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var video models.Video
		err := rows.Scan(
			&video.ID,
			&video.Author.Username,
			&video.Title,
			&video.Description,
			&video.CreatedAt,
			&video.Views,
			&video.Likes,
			&video.Comments,
			&video.Thumbnail,
			&video.Filepath,
		)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func (r *VideosRepo) AddComment(ctx context.Context, videoID int, comment models.Comment) error {
	query := `INSERT INTO comments (video_id, user_id, content, created_at , likes) VALUES ($1, $2, $3, $4 , $5)`

	_, err := r.db.Exec(ctx, query, videoID, comment.Author.Username, comment.Content, comment.CreatedAt, comment.Likes)
	if err != nil {
		return err
	}
	return nil
}

func (r *VideosRepo) LikeVideo(ctx context.Context, videoID int) error {
	query := `UPDATE videos SET likes = likes + 1 WHERE id=$1`

	_, err := r.db.Exec(ctx, query, videoID)
	if err != nil {
		return err
	}
	return nil
}

func (r *VideosRepo) UnlikeVideo(ctx context.Context, videoID int) error {
	query := `UPDATE videos SET likes = likes - 1 WHERE id=$1`

	_, err := r.db.Exec(ctx, query, videoID)
	if err != nil {
		return err
	}
	return nil
}
