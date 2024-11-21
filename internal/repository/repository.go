package repository

import (
	"context"
	"myTube/internal/models"

	"github.com/jackc/pgx/v5"
)

type Videos interface {
	GetById(ctx context.Context, id int) (models.Video, error)
	GetTrendy(ctx context.Context) ([]models.Video, error)
	Create(ctx context.Context , video models.Video) error
	Update(ctx context.Context, id int, video models.Video) error
	Delete(ctx context.Context, id int) error
	Search(ctx context.Context, query string) ([]models.Video, error)
	AddComment(ctx context.Context, videoID int, comment models.Comment) error
	LikeVideo(ctx context.Context, videoID int) error
	UnlikeVideo(ctx context.Context, videoID int) error
}

type Users interface {
	Create(user models.User) error
	GetByCredentials(ctx context.Context , email , password string) (models.User, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (models.User, error)
	GetByUsername(ctx context.Context, username string) (models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id int) error
}

type Repositories struct {
	Videos Videos
	Users Users
}

func NewRepositories(db *pgx.Conn) *Repositories {
	return &Repositories{
		Videos: NewVideoRepository(db),
		Users: NewUserRepository(db),
	}
}
