package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/app/controller"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/middleware"
)

func main() {
	envLoad()
	e := echo.New()
	conn := db.NewDB()
	e.Use(middleware.DBMiddlewareFunc(conn))
	controller.InitRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func envLoad() {
	if "" == os.Getenv("GO_ENV") {
		_ = os.Setenv("GO_ENV", "local")
	}
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatalf("Error loading env target env is %s", os.Getenv("GO_ENV"))
	}
}
