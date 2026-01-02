package service

import (
	"context"

	"github.com/nrbernard/yak-saver/internal/database"
)

type ProjectService struct {
	Repo *database.Queries
}

func NewProjectService(db *database.Queries) *ProjectService {
	return &ProjectService{Repo: db}
}

func (s *ProjectService) GetProjects(ctx context.Context) ([]database.Project, error) {
	return s.Repo.GetProjects(ctx)
}
