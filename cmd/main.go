package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/app/controller"
	"github.com/naoyakurokawa/ddd_menta/config"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/useruc"
)

func main() {
	userRepository := repoimpl.NewUserRepositoryImpl(config.NewDB())
	userUsecase := useruc.NewUserUsecase(userRepository)
	userHandler := controller.NewUserHandler(userUsecase)

	e := echo.New()
	controller.InitRouting(e, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
