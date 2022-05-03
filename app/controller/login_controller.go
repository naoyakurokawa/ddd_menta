package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/auth/usecase/loginuc"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
)

type loginRequest struct {
	Email    string
	Password string
}

func NewLoginController() echo.HandlerFunc {
	return func(c echo.Context) error {
		type res struct {
			Token string `json:"token"`
		}

		conn := db.NewDB()
		userRepository := repoimpl.NewUserRepositoryImpl(conn)
		loginUsecase := loginuc.NewLoginUsecase(userRepository)

		var req loginRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		t, err := loginUsecase.Login(req.Email, req.Password)

		if err != nil {
			log.Printf("failed to NewLoginController: %+v", err)
			return c.JSON(http.StatusBadRequest, "failed login")
		}

		response := res{
			Token: t,
		}

		return c.JSON(http.StatusOK, response)
	}
}
