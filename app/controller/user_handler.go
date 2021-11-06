package controller

import (
	"net/http"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/useruc"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/labstack/echo"
)

type UserHandler struct {
	userCreateUsecase useruc.UserCreateUsecase
}

// NewUserHandler user handlerのコンストラクタ
func NewUserHandler(uu useruc.UserCreateUsecase) *UserHandler {
	return &UserHandler{userCreateUsecase: uu}
}

type requestUser struct {
	Name     string
	Email    string
	Password string
	Profile  string
}

// Create userを保存するときのハンドラー
func UserCreate() echo.HandlerFunc {
	userRepository := repoimpl.NewUserRepositoryImpl(repoimpl.NewDB())
	userCreateUsecase := useruc.NewUserCreateUsecase(userRepository)
	return func(c echo.Context) error {
		type responseUser struct {
			userID userdm.UserID
		}
		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		//usecaseのCreate → infraのCreate
		createdUser, err := userCreateUsecase.Create(req.Name, req.Email, req.Password, req.Profile)
		if err != nil {
			return err
		}

		res := responseUser{
			userID: createdUser.UserID,
		}

		return c.JSON(http.StatusCreated, res)
	}
}