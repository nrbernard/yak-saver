package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nrbernard/yak-saver/data"
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

	e.GET("/projects", func(c echo.Context) error {
		projects := data.GetProjects()
		response := map[string]interface{}{
			"projects": projects,
		}
		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
