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

func (h *TaskHandler) CreateTask(c echo.Context) error {
	projectID := c.FormValue("project_id")
	parentTaskID := c.FormValue("parent_task_id")
	content := c.FormValue("content")
	link := c.FormValue("link")
	intProjectID, err := strconv.ParseInt(projectID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Project ID"})
	}
	intParentTaskID, err := strconv.ParseInt(parentTaskID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Parent Task ID"})
	}
	return h.TaskService.CreateTask(c.Request().Context(), intProjectID, sql.NullInt64{Int64: intParentTaskID, Valid: parentTaskID != ""}, content, sql.NullString{String: link, Valid: link != ""})
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	return h.TaskService.DeleteTask(c.Request().Context(), intID)
}
