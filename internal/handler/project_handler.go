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
