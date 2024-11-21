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