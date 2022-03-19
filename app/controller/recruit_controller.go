package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/recruituc"
)

type RecruitController struct {
	createRecruitUsecase recruituc.CreateRecruitUsecase
}

func NewRecruitController(ru recruituc.CreateRecruitUsecase) *RecruitController {
	return &RecruitController{createRecruitUsecase: ru}
}

type recruitRequest struct {
	Title         string
	Budget        uint32
	RecruitType   uint16
	Detail        string
	RecruitStatus uint16
}

func NewCreateRecruitController() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := db.NewDB()
		recruitRepository := repoimpl.NewRecruitRepositoryImpl(conn)
		recruitCreateUsecase := recruituc.NewCreateRecruitUsecase(recruitRepository)

		var req recruitRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		//TODO sessionからuserID取得するように変更
		userID := "e2e908dc-5981-4c4a-8e98-4487d3e122ad"
		err := recruitCreateUsecase.Create(
			userID,
			req.Title,
			req.Budget,
			req.RecruitType,
			req.Detail,
			req.RecruitStatus,
		)

		if err != nil {
			log.Printf("failed to NewCreateRecruitController: %+v", err)
			return c.JSON(http.StatusCreated, "failed create recruit")
		}

		return c.JSON(http.StatusCreated, "success create recruit")
	}
}
