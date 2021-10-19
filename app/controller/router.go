package controller

import (
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, userHandler UserHandler) {

	e.POST("/user", userHandler.Post())

}
