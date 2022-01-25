package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/app/controller"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/middleware"
)

func main() {
	e := echo.New()
	conn := db.NewDB()
	e.Use(middleware.DBMiddlewareFunc(conn))
	controller.InitRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}
