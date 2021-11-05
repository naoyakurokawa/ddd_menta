package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/app/controller"
)

func main() {
	e := echo.New()
	controller.InitRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}
