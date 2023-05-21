package main

import (
	"net/http"

	"github.com/claustra01/hackz-tsumaguro-websocket/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/socket", api.SocketHandler())
	e.GET("/close/:id", api.CloseHandler())

	e.Logger.Fatal(e.Start(":8501"))
}
