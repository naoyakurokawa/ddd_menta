package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/contractdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/contractuc"
)

type ContractController struct {
	createContractUsecase contractuc.CreateContractUsecase
}

func NewContractController(cu contractuc.CreateContractUsecase) *ContractController {
	return &ContractController{createContractUsecase: cu}
}

type contractRequest struct {
	ContractID string
	MentorID   string
	PlanID     string
}

func NewCreateContractController() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := db.NewDB()
		contractRepository := repoimpl.NewContractRepositoryImpl(conn)
		mentorRepository := repoimpl.NewMentorRepositoryImpl(conn)
		contractCreateUsecase := contractuc.NewCreateContractUsecase(contractRepository, mentorRepository)

		var req contractRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		//TODO sessionからuserID取得するように変更
		userID := "e2e908dc-5981-4c4a-8e98-4487d3e122ad"
		err := contractCreateUsecase.Create(
			userID,
			req.MentorID,
			req.PlanID,
		)

		if err != nil {
			log.Printf("failed to NewCreateContractController: %+v", err)
			return c.JSON(http.StatusBadRequest, "failed create contract")
		}

		return c.JSON(http.StatusCreated, "success create contract")
	}
}

func NewUpdateUnderContractController() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := db.NewDB()
		contractRepository := repoimpl.NewContractRepositoryImpl(conn)
		updateUnderContractUsecase := contractuc.NewUpdateContractStatusUsecase(contractRepository)

		var req contractRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err := updateUnderContractUsecase.UpdateContractStatus(req.ContractID, contractdm.UnderContract.Uint16())
		if err != nil {
			log.Printf("failed to NewUpdateUnderContractController: %+v", err)
			return c.JSON(http.StatusBadRequest, "can't update contract")
		}

		return c.JSON(http.StatusCreated, "success update contract")
	}
}

func NewUpdateTerminatedContractController() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := db.NewDB()
		contractRepository := repoimpl.NewContractRepositoryImpl(conn)
		updateUnderContractUsecase := contractuc.NewUpdateContractStatusUsecase(contractRepository)

		var req contractRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err := updateUnderContractUsecase.UpdateContractStatus(req.ContractID, contractdm.TerminatedContract.Uint16())
		if err != nil {
			log.Printf("failed to NewUpdateUnderContractController: %+v", err)
			return c.JSON(http.StatusCreated, "can't update contract")
		}

		return c.JSON(http.StatusCreated, "success update contract")
	}
}
