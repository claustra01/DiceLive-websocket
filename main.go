package main

import (
	"net/http"

	"github.com/claustra01/hackz-tsumaguro-websocket/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewCors() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {
			c.Response().Writer.Header().Set("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))
			c.Response().Header().Set("Access-Control-Max-Age", "12h0m0s")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Authorization")
			c.Response().Header().Set("Access-Control-Expose-Headers", "Content-Length")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			if c.Request().Method == http.MethodOptions {
				return c.NoContent(http.StatusNoContent)
			}

			return next(c)
		}
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(NewCors())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/socket", api.SocketHandler())
	e.GET("/close/:id", api.CloseHandler())

	e.Logger.Fatal(e.Start(":8501"))
}
