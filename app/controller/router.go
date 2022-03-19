package controller

import (
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo) {

	e.POST("/user/create", NewCreateUserController())
	e.POST("/mentor/create", NewCreateMentorController())
	e.POST("/contract/create", NewCreateContractController())
	e.POST("/contract/update_under_contract", NewUpdateUnderContractController())
	e.POST("/contract/update_terminated_contract", NewUpdateTerminatedContractController())
	e.POST("/recruit/create", NewCreateRecruitController())
	e.POST("/suggestion/create", NewCreateSuggestionController())

}
