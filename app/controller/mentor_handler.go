package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
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

type requestMentor struct {
	Title                 string
	MainImg               string
	SubImg                string
	Category              string
	Detial                string
	MentorTag             []string
	MentorAssessment      []string
	MentorExperienceYears []string
	PlanTitle             []string
	PlanCategory          []string
	PlanTag               []string
	PlanDetial            []string
	PlanType              []string
	PlanPrice             []string
	PlanStatus            []string
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
		var req requestMentor
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		//TODO sessionからuserID取得するように変更
		userID := "e2e908dc-5981-4c4a-8e98-4487d3e122ad"
		//usecaseのCreate → infraのCreate
		createdMentor, err := mentorCreateUsecase.Create(
			userID,
			req.Title,
			req.MainImg,
			req.SubImg,
			req.Category,
			req.Detial,
			req.MentorTag,
			req.MentorAssessment,
			req.MentorExperienceYears,
			req.PlanTitle,
			req.PlanCategory,
			req.PlanTag,
			req.PlanDetial,
			req.PlanType,
			req.PlanPrice,
			req.PlanStatus,
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
