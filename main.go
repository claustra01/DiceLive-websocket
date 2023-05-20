package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()            // インスタンスを作成
	e.Use(middleware.Logger()) // ミドルウェアを設定

	e.GET("/", func(c echo.Context) error { // ルートを設定
		return c.String(http.StatusOK, "Hello, World!") // 出力
	})

	e.Logger.Fatal(e.Start(":8001")) // サーバーをポート番号8001で起動
}
