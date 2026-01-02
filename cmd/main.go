package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nrbernard/yak-saver/data"
)

func main() {
	e := echo.New()

	e.GET("/projects", func(c echo.Context) error {
		projects := data.GetProjects()
		response := map[string]interface{}{
			"projects": projects,
		}
		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
