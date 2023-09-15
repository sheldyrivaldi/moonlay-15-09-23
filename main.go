package main

import (
	"log"
	"moonlay-todolist/database"
	"moonlay-todolist/database/migration"
	"moonlay-todolist/internal/factory"
	"moonlay-todolist/internal/http"
	"moonlay-todolist/internal/pkg/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Init()
	migration.Init()

	e := echo.New()

	middleware.Init(e)
	f := factory.NewFactory()
	http.Init(e, f)

	e.Logger.Fatal(e.Start(":5000"))
}
