package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/useruc"
)

type UserController struct {
	userCreateUsecase useruc.UserCreateUsecase
}

func NewUserController(uu useruc.UserCreateUsecase) *UserController {
	return &UserController{userCreateUsecase: uu}
}

type requestUser struct {
	Name            string
	Email           string
	Password        string
	Profile         string
	From            []string
	To              []string
	Detail          []string
	Tag             []string
	Assessment      []string
	ExperienceYears []string
}

// Create userを保存するときのハンドラー
func NewCreateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := db.NewDB()
		userRepository := repoimpl.NewUserRepositoryImpl(conn)
		userCreateUsecase := useruc.NewUserCreateUsecase(userRepository)
		type responseUser struct {
			userID userdm.UserID
		}
		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		//usecaseのCreate → infraのCreate
		createdUser, err := userCreateUsecase.Create(req.Name, req.Email, req.Password, req.Profile, req.From, req.To, req.Detail, req.Tag, req.Assessment, req.ExperienceYears)
		if err != nil {
			return err
		}

		res := responseUser{
			userID: createdUser.UserID(),
		}

		return c.JSON(http.StatusCreated, res)
	}
}