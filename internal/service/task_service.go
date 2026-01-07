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

func (s *TaskService) GetTask(ctx context.Context, id int64) (database.Task, error) {
	return s.Repo.GetTask(ctx, id)
}

func (s *TaskService) UpdateTask(ctx context.Context, id int64, content string, link sql.NullString, completedAt sql.NullTime) error {
	return s.Repo.UpdateTask(ctx, database.UpdateTaskParams{
		ID:          id,
		Content:     content,
		Link:        link,
		CompletedAt: completedAt,
	})
}

func (s *TaskService) CreateTask(ctx context.Context, projectID int64, parentTaskID sql.NullInt64, content string, link sql.NullString) (database.Task, error) {
	return s.Repo.CreateTask(ctx, database.CreateTaskParams{
		ProjectID:    projectID,
		ParentTaskID: parentTaskID,
		Content:      content,
		Link:         link,
	})
}

// collectDescendantIDs recursively collects all descendant task IDs
func (s *TaskService) collectDescendantIDs(ctx context.Context, parentID int64) ([]int64, error) {
	var allIDs []int64

	// Get direct children
	children, err := s.Repo.GetTasksByParentID(ctx, sql.NullInt64{Int64: parentID, Valid: true})
	if err != nil {
		return nil, err
	}

	// For each child, collect its ID and recursively collect its descendants
	for _, child := range children {
		allIDs = append(allIDs, child.ID)
		descendants, err := s.collectDescendantIDs(ctx, child.ID)
		if err != nil {
			return nil, err
		}
		allIDs = append(allIDs, descendants...)
	}

	return allIDs, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int64) error {
	// Collect all descendant IDs
	descendantIDs, err := s.collectDescendantIDs(ctx, id)
	if err != nil {
		return err
	}

	// Delete all descendants first (from deepest to shallowest doesn't matter since we're deleting by ID)
	for _, descendantID := range descendantIDs {
		if err := s.Repo.DeleteTask(ctx, descendantID); err != nil {
			return err
		}
	}

	// Finally, delete the task itself
	return s.Repo.DeleteTask(ctx, id)
}
