package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nrbernard/yak-saver/internal/database"
	"github.com/nrbernard/yak-saver/internal/handler"
	"github.com/nrbernard/yak-saver/internal/service"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	db, err := sql.Open("sqlite3", "data/yak-saver.db")
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
	e.PUT("/tasks/:id", taskHandler.UpdateTask)
	e.POST("/tasks", taskHandler.CreateTask)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}
