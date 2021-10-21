package controller

import (
	"net/http"

	"github.com/naoyakurokawa/ddd_menta/core/usecase/useruc"

	"github.com/labstack/echo"
)

// UserHandler user handlerのinterface
type UserHandler interface {
	Post() echo.HandlerFunc
}

type UserHandlerImpl struct {
	userUsecase useruc.UserUsecase
}

// NewUserHandler user handlerのコンストラクタ
func NewUserHandler(userUsecase useruc.UserUsecase) UserHandler {
	return &UserHandlerImpl{userUsecase: userUsecase}
}

type requestUser struct {
	Name     string
	Email    string
	Password string
	Profile  string
}

type responseUser struct {
	UserId string
}

// Post userを保存するときのハンドラー
func (uh *UserHandlerImpl) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		//usecaseのCreate　→ infraのCreate
		createdUser, err := uh.userUsecase.Create(req.Name, req.Email, req.Password, req.Profile)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			UserId: createdUser.UserId,
		}

		return c.JSON(http.StatusCreated, res)
	}
}
