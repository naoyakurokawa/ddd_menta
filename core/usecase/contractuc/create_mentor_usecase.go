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
		status string,
	) error
}

type CreateContractUsecaseImpl struct {
	contractRepo contractdm.ContractRepository
}

// user usecaseのコンストラクタ
func NewCreateContractUsecase(contractRepo contractdm.ContractRepository) CreateContractUsecase {
	return &CreateContractUsecaseImpl{contractRepo: contractRepo}
}

// Create userを保存するときのユースケース
func (cu *CreateContractUsecaseImpl) Create(
	userID string,
	planID string,
	status string,
) error {
	// contractID := contractdm.NewContractID()
	userIDIns, err := userdm.NewUserIDByVal(userID)
	if err != nil {
		return err
	}
	planIDIns, err := mentordm.NewPlanIDByVal(planID)
	if err != nil {
		return err
	}
	castedStatus, err := contractdm.StrCastUint(status)
	if err != nil {
		return err
	}
	statusIns, err := contractdm.NewStatus(castedStatus)
	if err != nil {
		return err
	}

	// メンター作成
	contract, err := contractdm.NewContract(
		userIDIns,
		planIDIns,
		statusIns,
	)
	if err != nil {
		return err
	}

	//最終的にinfraのCreateメソッドを実行することになる
	err = cu.contractRepo.Create(contract)
	if err != nil {
		return err
	}

	return nil
}
