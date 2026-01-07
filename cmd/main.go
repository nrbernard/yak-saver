package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nrbernard/yak-saver/internal/database"
	"github.com/nrbernard/yak-saver/internal/handler"
	"github.com/nrbernard/yak-saver/internal/service"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	e := echo.New()

	// CORS is only needed for local development when frontend and backend are on different ports
	corsOrigins := getEnv("CORS_ORIGINS", "http://localhost:5173")
	if corsOrigins != "" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: strings.Split(corsOrigins, ","),
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.PATCH, echo.DELETE},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		}))
	}

	db, err := sql.Open("sqlite3", getEnv("DATABASE_PATH", "data/yak-saver.db"))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Successfully connected to database")
	dbQueries := database.New(db)

	projectService := service.NewProjectService(dbQueries)
	taskService := service.NewTaskService(dbQueries)

	projectHandler := handler.NewProjectHandler(projectService)
	taskHandler := handler.NewTaskHandler(taskService)

	e.GET("/projects", projectHandler.GetProjects)
	e.POST("/projects", projectHandler.CreateProject)
	e.DELETE("/projects/:id", projectHandler.DeleteProject)
	e.PATCH("/tasks/:id", taskHandler.UpdateTask)
	e.POST("/tasks", taskHandler.CreateTask)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask)

	// Serve frontend static files if configured
	staticPath := getEnv("STATIC_FILES_PATH", "")
	if staticPath != "" {
		// Serve static assets
		e.Static("/assets", filepath.Join(staticPath, "assets"))
		e.File("/vite.svg", filepath.Join(staticPath, "vite.svg"))

		// SPA fallback - must come last
		e.File("/*", filepath.Join(staticPath, "index.html"))
	}

	e.Logger.Fatal(e.Start(":" + getEnv("PORT", "8080")))
}
