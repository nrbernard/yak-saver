package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nrbernard/yak-saver/internal/service"
)

type TaskHandler struct {
	TaskService *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{TaskService: service}
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	content := c.FormValue("content")
	link := c.FormValue("link")
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	return h.TaskService.UpdateTask(c.Request().Context(), intID, content, sql.NullString{String: link, Valid: link != ""})
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
