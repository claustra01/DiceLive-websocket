package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SocketHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// ハンドラの実装
		return c.String(http.StatusOK, "WebSocket")
	}
}
