package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/suggestionuc"
)

type SuggestionController struct {
	createSuggestionUsecase suggestionuc.CreateSuggestionUsecase
}

func NewSuggestionController(su suggestionuc.CreateSuggestionUsecase) *SuggestionController {
	return &SuggestionController{createSuggestionUsecase: su}
}

type suggestionRequest struct {
	RecruitID        string
	Price            uint32
	SuggestionType   uint16
	Detail           string
	SuggestionStatus uint16
}

func NewCreateSuggestionController() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := db.NewDB()
		suggestionRepository := repoimpl.NewSuggestionRepositoryImpl(conn)
		mentorRepository := repoimpl.NewMentorRepositoryImpl(conn)
		recruitRepository := repoimpl.NewRecruitRepositoryImpl(conn)
		suggestionCreateUsecase := suggestionuc.NewCreateSuggestionUsecase(
			suggestionRepository,
			mentorRepository,
			recruitRepository,
		)

		var req suggestionRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		//TODO sessionからuserID取得するように変更
		mentorID := "e2e908dc-5981-4c4a-8e98-4487d3e122ad"
		err := suggestionCreateUsecase.Create(
			mentorID,
			req.RecruitID,
			req.Price,
			req.SuggestionType,
			req.Detail,
			req.SuggestionStatus,
		)

		if err != nil {
			log.Printf("failed to NewCreateSuggestionController: %+v", err)
			return c.JSON(http.StatusCreated, "failed create suggestion")
		}

		return c.JSON(http.StatusCreated, "success create suggestion")
	}
}
