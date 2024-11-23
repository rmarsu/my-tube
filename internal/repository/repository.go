package repository

import (
	"context"
	"myTube/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Videos interface {
	GetById(ctx context.Context, id int) (models.Video, error)
	GetTrendy(ctx context.Context) ([]models.Video, error)
	Create(ctx context.Context, video models.Video) error
	Update(ctx context.Context, id int, video models.Video) error
	Delete(ctx context.Context, id int) error
	Search(ctx context.Context, query string) ([]models.Video, error)
	LikeVideo(ctx context.Context, videoID int, userID int) error
	UnlikeVideo(ctx context.Context, videoID int, userID int) error
}

type Users interface {
	Create(user models.User) error
	GetByCredentials(ctx context.Context, email, password string) (models.User, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (models.User, error)
	GetByUsername(ctx context.Context, username string) (models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id int) error
}

type Repositories struct {
	Videos Videos
	Users  Users
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		Videos: NewVideoRepository(db),
		Users:  NewUserRepository(db),
	}
}
