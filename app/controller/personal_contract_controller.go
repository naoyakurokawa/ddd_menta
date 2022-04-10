package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/personalcontractuc"
)

type PersonalContractController struct {
	createPersonalContractUsecase personalcontractuc.CreatePersonalContractUsecase
}

func NewPersonalContractController(cu personalcontractuc.CreatePersonalContractUsecase) *PersonalContractController {
	return &PersonalContractController{createPersonalContractUsecase: cu}
}

type personalContractRequest struct {
	SuggestionID string
}

func NewCreatePersonalContractController() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn := db.NewDB()
		personalContractRepository := repoimpl.NewPersonalContractRepositoryImpl(conn)
		contractCreateUsecase := personalcontractuc.NewCreatePersonalContractUsecase(personalContractRepository)

		var req personalContractRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err := contractCreateUsecase.Create(
			req.SuggestionID,
		)

		if err != nil {
			log.Printf("failed to NewCreatePersonalContractController: %+v", err)
			return c.JSON(http.StatusBadRequest, "failed create personal_contract")
		}

		return c.JSON(http.StatusCreated, "success create personal_contract")
	}
}
