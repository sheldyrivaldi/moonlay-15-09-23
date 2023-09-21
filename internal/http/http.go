package http

import (
	"moonlay-todolist/internal/app/list"
	"moonlay-todolist/internal/app/sublist"
	"moonlay-todolist/internal/factory"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f *factory.Factory) {

	// index
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// routes api
	list.NewHandler(f).Route(e.Group("/list"))
	sublist.NewHandler(f).Route((e.Group("/list/:listID/sub")))
}
