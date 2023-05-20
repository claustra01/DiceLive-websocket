package api

import (
	"net/http"

	"github.com/claustra01/hackz-tsumaguro-websocket/common"
	"github.com/labstack/echo/v4"
)

func CloseHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		delete(common.SocketList, c.Param("id"))
		delete(common.CommentList, c.Param("id"))
		delete(common.ReactionList, c.Param("id"))
		return c.String(http.StatusOK, "Stream Closed!")
	}
}
