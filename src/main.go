package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		errorResponse := GetNotFoundErrorResponse("page", "چیزی")
		return c.JSON(http.StatusNotFound, errorResponse)
	}
	e := echo.New()

	//v1 := e.Group("/api/v1")
	e.Logger.Fatal(e.Start(":80"))
}
