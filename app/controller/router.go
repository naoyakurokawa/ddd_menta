package controller

import (
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo) {

	e.POST("/user_create", NewCreateUserController())
	e.POST("/mentor_create", NewCreateMentorController())

}
