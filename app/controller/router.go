package controller

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/naoyakurokawa/ddd_menta/auth/infrastructure/jwt"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo) {
	r := e.Group("/contract")
	r.Use(middleware.JWTWithConfig(jwt.JwtConfig))
	e.POST("/user/create", NewCreateUserController())
	e.POST("/mentor/create", NewCreateMentorController())
	r.POST("/create", NewCreateContractController())
	r.POST("/update_under_contract", NewUpdateUnderContractController())
	r.POST("/update_terminated_contract", NewUpdateTerminatedContractController())
	e.POST("/recruit/create", NewCreateRecruitController())
	e.POST("/suggestion/create", NewCreateSuggestionController())
	e.POST("/personal_contract/create", NewCreatePersonalContractController())
	e.POST("/login", NewLoginController())
}
