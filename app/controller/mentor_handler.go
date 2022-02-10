package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/mentoruc"
)

type MentorHandler struct {
	mentorCreateUsecase mentoruc.MentorCreateUsecase
}

// NewUserHandler user handlerのコンストラクタ
func NewMentorHandler(mu mentoruc.MentorCreateUsecase) *MentorHandler {
	return &MentorHandler{mentorCreateUsecase: mu}
}

type requestmentor struct {
	title                 string
	mainImg               string
	subImg                string
	category              string
	detial                string
	mentorTag             []string
	mentorAssessment      []string
	mentorExperienceYears []string
	planTitle             []string
	planCategory          []string
	planTag               []string
	planDetial            []string
	planType              []string
	planPrice             []string
	planStatus            []string
}

// Create userを保存するときのハンドラー
func MentorCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := db.NewDB()
		mentorRepository := repoimpl.NewMentorRepositoryImpl(conn)
		mentorCreateUsecase := mentoruc.NewMentorCreateUsecase(mentorRepository)
		type responseMentor struct {
			mentorID mentordm.MentorID
		}
		var req requestmentor
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		//TODO sessionからuserID取得するように変更
		userID, err := userdm.NewUserID()
		if err != nil {
			return err
		}
		//usecaseのCreate → infraのCreate
		createdMentor, err := mentorCreateUsecase.Create(
			userID,
			req.title,
			req.mainImg,
			req.subImg,
			req.category,
			req.detial,
			req.mentorTag,
			req.mentorAssessment,
			req.mentorExperienceYears,
			req.planTitle,
			req.planCategory,
			req.planTag,
			req.planDetial,
			req.planType,
			req.planPrice,
			req.planStatus,
		)

		if err != nil {
			return err
		}

		res := responseMentor{
			mentorID: createdMentor.MentorID(),
		}

		return c.JSON(http.StatusCreated, res)
	}
}
