package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/mentoruc"
)

type MentorController struct {
	createMentorUsecase mentoruc.CreateMentorUsecase
}

func NewMentorController(mu mentoruc.CreateMentorUsecase) *MentorController {
	return &MentorController{createMentorUsecase: mu}
}

type mentorRequest struct {
	Title                 string
	MainImg               string
	SubImg                string
	Category              string
	Detail                string
	MentorTag             []string
	MentorAssessment      []string
	MentorExperienceYears []string
	Plans                 []mentoruc.Plan
}

func NewCreateMentorController() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := db.NewDB()
		mentorRepository := repoimpl.NewMentorRepositoryImpl(conn)
		mentorCreateUsecase := mentoruc.NewCreateMentorUsecase(mentorRepository)

		var req mentorRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		//TODO sessionからuserID取得するように変更
		userID := "e2e908dc-5981-4c4a-8e98-4487d3e122ad"
		err := mentorCreateUsecase.Create(
			userID,
			req.Title,
			req.MainImg,
			req.SubImg,
			req.Category,
			req.Detail,
			req.MentorTag,
			req.MentorAssessment,
			req.MentorExperienceYears,
			req.Plans,
		)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, "success create mentor")
	}
}
