package contractuc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/contractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

type CreateContractUsecase interface {
	Create(
		userID string,
		planID string,
	) error
}

type CreateContractUsecaseImpl struct {
	contractRepo contractdm.ContractRepository
}

func NewCreateContractUsecase(contractRepo contractdm.ContractRepository) CreateContractUsecase {
	return &CreateContractUsecaseImpl{contractRepo: contractRepo}
}

// Create userを保存するときのユースケース
func (cu *CreateContractUsecaseImpl) Create(
	userID string,
	planID string,
) error {
	userIDIns, err := userdm.NewUserIDByVal(userID)
	if err != nil {
		return err
	}
	planIDIns, err := mentordm.NewPlanIDByVal(planID)
	if err != nil {
		return err
	}
	//メンティーによる契約リクエスト時のStatusは未承認
	status := contractdm.Unapproved

	contract, err := contractdm.NewContract(
		userIDIns,
		planIDIns,
		status,
	)
	if err != nil {
		return err
	}

	err = cu.contractRepo.Create(contract)
	if err != nil {
		return err
	}

	return nil
}
