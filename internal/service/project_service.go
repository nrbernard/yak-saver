package service

import (
	"context"
	"strconv"

	"github.com/nrbernard/yak-saver/internal/database"
)

type ProjectService struct {
	Repo *database.Queries
}

func NewProjectService(db *database.Queries) *ProjectService {
	return &ProjectService{Repo: db}
}

func (s *ProjectService) GetProjects(ctx context.Context) ([]map[string]interface{}, error) {
	projects, err := s.Repo.GetProjects(ctx)
	if err != nil {
		return nil, err
	}

	tasks, err := s.Repo.GetTasksOrdered(ctx)
	if err != nil {
		return nil, err
	}

	// Build maps for quick lookup
	projectMap := make(map[int64]map[string]interface{})
	taskNodes := make(map[int64]map[string]interface{})
	projectTasks := make(map[int64][]map[string]interface{})

	// Initialize project structures
	for _, project := range projects {
		projectMap[project.ID] = map[string]interface{}{
			"id":    strconv.FormatInt(project.ID, 10),
			"name":  project.Name,
			"tasks": []map[string]interface{}{},
		}
		projectTasks[project.ID] = []map[string]interface{}{}
	}

	// Build task tree in a single pass (tasks are ordered: top-level first)
	for _, task := range tasks {
		taskNode := map[string]interface{}{
			"id":       strconv.FormatInt(task.ID, 10),
			"content":  task.Content,
			"children": []map[string]interface{}{},
		}
		if task.Link.Valid {
			taskNode["link"] = task.Link.String
		}

		taskNodes[task.ID] = taskNode

		if !task.ParentTaskID.Valid {
			// Top-level task: add to project
			projectTasks[task.ProjectID] = append(projectTasks[task.ProjectID], taskNode)
		} else {
			// Child task: add to parent's children
			parent := taskNodes[task.ParentTaskID.Int64]
			children := parent["children"].([]map[string]interface{})
			parent["children"] = append(children, taskNode)
		}
	}

	// Attach tasks to projects
	result := make([]map[string]interface{}, 0, len(projects))
	for _, project := range projects {
		proj := projectMap[project.ID]
		proj["tasks"] = projectTasks[project.ID]
		result = append(result, proj)
	}

	return result, nil
}
