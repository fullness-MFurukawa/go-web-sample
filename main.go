package main

import (
	handler "go-web-sample/presentation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンス生成
	e := echo.New()
	// ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Debug = true

	// ルートを設定
	e.GET("/search", handler.Search)
	e.POST("/register", handler.Register)
	// ポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}
