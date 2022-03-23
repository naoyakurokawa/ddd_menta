package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/useruc"
)

type UserController struct {
	userCreateUsecase useruc.CreateUserUsecase
}

func NewUserController(uu useruc.CreateUserUsecase) *UserController {
	return &UserController{userCreateUsecase: uu}
}

type requestUser struct {
	Name        string
	Email       string
	Password    string
	Profile     string
	UserCareers []useruc.UserCareer
	UserSkills  []useruc.UserSkill
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
		err := userCreateUsecase.Create(
			req.Name,
			req.Email,
			req.Password,
			req.Profile,
			req.UserCareers,
			req.UserSkills,
		)

		if err != nil {
			log.Printf("failed to NewCreateUserController: %+v", err)
			return c.JSON(http.StatusCreated, "failed create user")
		}

		return c.JSON(http.StatusCreated, "success create user")
	}
}
