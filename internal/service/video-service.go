package service

import (
	"context"
	"myTube/internal/models"
	"myTube/internal/repository"
)


type VideoService struct {
	repo repository.Videos
}

func NewVideoService(repo repository.Videos) *VideoService {
     return &VideoService{repo: repo}
}

func (s *VideoService) GetVideo(ctx context.Context, id int) (models.Video, error) {
     return s.repo.GetById(ctx ,id)
}

func (s *VideoService) GetTrendyVideos(ctx context.Context) ([]models.Video, error) {
     return s.repo.GetTrendy(ctx)
}

func (s *VideoService) CreateVideo(ctx context.Context ,video models.Video) error {
     return s.repo.Create(ctx ,video)
}

func (s *VideoService) UpdateVideo(ctx context.Context, id int, video models.Video) error {
     return s.repo.Update(ctx, id, video)
}

func (s *VideoService) DeleteVideo(ctx context.Context, id int) error {
     return s.repo.Delete(ctx, id)
}

func (s *VideoService) SearchVideos(ctx context.Context, query string) ([]models.Video, error) {
     return s.repo.Search(ctx, query)
}

func (s *VideoService) LikeVideo(ctx context.Context, videoID int, userID int) error {
     return s.repo.LikeVideo(ctx, videoID, userID)
}

func (s *VideoService) UnlikeVideo(ctx context.Context, videoID int, userID int) error {
     return s.repo.UnlikeVideo(ctx, videoID, userID)
}

