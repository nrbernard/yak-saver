package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nrbernard/yak-saver/internal/service"
)

type ProjectHandler struct {
	ProjectService *service.ProjectService
}

func NewProjectHandler(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{ProjectService: service}
}

func (h *ProjectHandler) GetProjects(c echo.Context) error {
	projects, err := h.ProjectService.GetProjects(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, projects)
}

type CreateProjectRequest struct {
	Name string `json:"name"`
}

func (h *ProjectHandler) CreateProject(c echo.Context) error {
	var req CreateProjectRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON: " + err.Error()})
	}

	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name is required"})
	}

	project, err := h.ProjectService.CreateProject(c.Request().Context(), req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create project: " + err.Error()})
	}

	projectResponse := map[string]interface{}{
		"id":    project.ID,
		"name":  project.Name,
		"tasks": []map[string]interface{}{},
	}

	return c.JSON(http.StatusCreated, projectResponse)
}
