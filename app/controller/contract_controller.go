package controller

import (
	"net/http"

	"github.com/labstack/echo"
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
	PlanID string
	Status string
}

func NewCreateContractController() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := db.NewDB()
		contractRepository := repoimpl.NewContractRepositoryImpl(conn)
		contractCreateUsecase := contractuc.NewCreateContractUsecase(contractRepository)

		var req contractRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		//TODO sessionからuserID取得するように変更
		userID := "e2e908dc-5981-4c4a-8e98-4487d3e122ad"
		err := contractCreateUsecase.Create(
			userID,
			req.PlanID,
			req.Status,
		)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, "success create contract")
	}
}
