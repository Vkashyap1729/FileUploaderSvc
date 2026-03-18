package service

import (
	"context"
	"file-uploader/internal/repository"
)

type UploadService struct {
	repo *repository.UploadRepository
}

func NewUploadService(repo *repository.UploadRepository) *UploadService {
	return &UploadService{repo: repo}
}

func (s *UploadService) GetAllUploads(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetAllUploads(ctx)
}