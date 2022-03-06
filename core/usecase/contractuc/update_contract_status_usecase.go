package contractuc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/contractdm"
	"golang.org/x/xerrors"
)

type UpdateContractStatusUsecase interface {
	UpdateContractStatus(
		contractID string,
		contractStatus uint16,
	) error
}

type UpdateContractStatusUsecaseImpl struct {
	contractRepo contractdm.ContractRepository
}

func NewUpdateContractStatusUsecase(
	contractRepo contractdm.ContractRepository,
) UpdateContractStatusUsecase {
	return &UpdateContractStatusUsecaseImpl{
		contractRepo: contractRepo,
	}
}

func (uc *UpdateContractStatusUsecaseImpl) UpdateContractStatus(
	contractID string,
	contractStatus uint16,
) error {
	contractIDIns, err := contractdm.NewContractIDByVal(contractID)
	if err != nil {
		return err
	}

	contract, err := uc.contractRepo.FindByID(contractIDIns)
	if err != nil {
		return err
	}
	contractStatusIns, err := contractdm.NewContractStatus(contractStatus)
	if err != nil {
		return err
	}
	if !contract.CanUpdateContractStatus(contractStatusIns) {
		return xerrors.New("can't update contract")
	}

	return uc.contractRepo.UpdateContractStatus(contractIDIns, contractStatusIns)
}
