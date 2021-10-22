package controller

import (
	"net/http"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/useruc"
	"golang.org/x/xerrors"

	"github.com/labstack/echo"
)

type UserHandler struct {
	userUsecase useruc.UserUsecase
}

// NewUserHandler user handlerのコンストラクタ
func NewUserHandler(uu useruc.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: uu}
}

type requestUser struct {
	Name     string
	Email    string
	Password string
	Profile  string
}

// Create userを保存するときのハンドラー
func (uh *UserHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		type responseUser struct {
			UserId userdm.UserId
		}
		var req requestUser
		if err := c.Bind(&req); err != nil {
			return xerrors.New("Insufficient request parameters")
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
