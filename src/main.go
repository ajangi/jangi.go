package main

import (
	"github.com/ajangi/jangi.go/src"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		errorResponse := src.GetNotFoundErrorResponse("page", "چیزی")
		return c.JSON(http.StatusNotFound, errorResponse)
	}
	e := echo.New()

	v1 := e.Group("/api/v1")
	v1.GET("/me",)
	e.Logger.Fatal(e.Start(":8080"))
}
