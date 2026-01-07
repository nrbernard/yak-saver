package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nrbernard/yak-saver/internal/service"
)

type TaskHandler struct {
	TaskService *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{TaskService: service}
}

type UpdateTaskRequest struct {
	Content   *string `json:"content"`
	Link      *string `json:"link"`
	Completed *bool   `json:"completed"`
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var req UpdateTaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON: " + err.Error()})
	}

	// Fetch current task to get existing values for omitted fields
	currentTask, err := h.TaskService.GetTask(c.Request().Context(), intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	// Use provided content or fall back to existing
	content := currentTask.Content
	if req.Content != nil {
		content = *req.Content
	}

	// Use provided link or fall back to existing
	link := currentTask.Link
	if req.Link != nil {
		link = sql.NullString{String: *req.Link, Valid: *req.Link != ""}
	}

	// Use provided completed or fall back to existing
	completedAt := currentTask.CompletedAt
	if req.Completed != nil {
		if *req.Completed {
			// Set to current server time when marking as complete
			completedAt = sql.NullTime{Time: time.Now(), Valid: true}
		} else {
			// Explicitly set to NULL when marking as incomplete
			completedAt = sql.NullTime{Valid: false}
		}
	}

	return h.TaskService.UpdateTask(c.Request().Context(), intID, content, link, completedAt)
}

type CreateTaskRequest struct {
	ProjectID    int64  `json:"projectId"`
	ParentTaskID *int64 `json:"parentTaskId,omitempty"`
	Content      string `json:"content"`
	Link         string `json:"link,omitempty"`
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var req CreateTaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON: " + err.Error()})
	}

	var parentTaskIDNull sql.NullInt64
	if req.ParentTaskID != nil {
		parentTaskIDNull = sql.NullInt64{Int64: *req.ParentTaskID, Valid: true}
	}

	task, err := h.TaskService.CreateTask(c.Request().Context(), req.ProjectID, parentTaskIDNull, req.Content, sql.NullString{String: req.Link, Valid: req.Link != ""})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create task: " + err.Error()})
	}

	taskNode := map[string]interface{}{
		"id":       task.ID,
		"content":  task.Content,
		"children": []map[string]interface{}{},
	}
	if task.Link.Valid {
		taskNode["link"] = task.Link.String
	}
	if task.CompletedAt.Valid {
		taskNode["completedAt"] = task.CompletedAt.Time
	}

	return c.JSON(http.StatusCreated, taskNode)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	return h.TaskService.DeleteTask(c.Request().Context(), intID)
}
