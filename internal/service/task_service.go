package service

import (
	"context"
	"database/sql"

	"github.com/nrbernard/yak-saver/internal/database"
)

type TaskService struct {
	Repo *database.Queries
}

func NewTaskService(db *database.Queries) *TaskService {
	return &TaskService{Repo: db}
}

func (s *TaskService) UpdateTask(ctx context.Context, id int64, content string, link sql.NullString) error {
	return s.Repo.UpdateTask(ctx, database.UpdateTaskParams{
		ID:      id,
		Content: content,
		Link:    link,
	})
}

func (s *TaskService) CreateTask(ctx context.Context, projectID int64, parentTaskID sql.NullInt64, content string, link sql.NullString) error {
	return s.Repo.CreateTask(ctx, database.CreateTaskParams{
		ProjectID:    projectID,
		ParentTaskID: parentTaskID,
		Content:      content,
		Link:         link,
	})
}

func (s *TaskService) DeleteTask(ctx context.Context, id int64) error {
	return s.Repo.DeleteTask(ctx, id)
}
