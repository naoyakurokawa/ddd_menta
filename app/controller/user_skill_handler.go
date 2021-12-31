package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/userskilluc"
)

type UserSkillHandler struct {
	userSkillCreateUsecase userskilluc.UserSkillCreateUsecase
}

// NewUserHandler user handlerのコンストラクタ
func NewUserSkillHandler(uu userskilluc.UserSkillCreateUsecase) *UserSkillHandler {
	return &UserSkillHandler{userSkillCreateUsecase: uu}
}

// TODO:UserIDはcookieから取得するように変更
type requestUserSkill struct {
	UserID          string
	Tag             []string
	Assessment      []string
	ExperienceYears []string
}

// Create userを保存するときのハンドラー
func UserSkillCreate() echo.HandlerFunc {
	userSkillRepository := repoimpl.NewUserSkillRepositoryImpl(repoimpl.NewDB())
	userSkillCreateUsecase := userskilluc.NewUserSkillCreateUsecase(userSkillRepository)
	return func(c echo.Context) error {
		type responseUser struct {
			userID userdm.UserID
		}
		var req requestUserSkill
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		//usecaseのCreate → infraのCreate
		createdUserSkill, err := userSkillCreateUsecase.Create(req.UserID, req.Tag, req.Assessment, req.ExperienceYears)
		if err != nil {
			return err
		}

		res := responseUser{
			userID: createdUserSkill[0].UserID(),
		}

		return c.JSON(http.StatusCreated, res)
	}
}
