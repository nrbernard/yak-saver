package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nrbernard/yak-saver/internal/database"
	"github.com/nrbernard/yak-saver/internal/handler"
	"github.com/nrbernard/yak-saver/internal/service"
)

func main() {
	e := echo.New()

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

	projectHandler := handler.NewProjectHandler(projectService)

	e.GET("/projects", projectHandler.GetProjects)

	e.Logger.Fatal(e.Start(":8080"))
}
