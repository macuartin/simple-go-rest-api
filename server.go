package main

import (
	"net/http"

	"gopkg.in/labstack/echo.v4"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
