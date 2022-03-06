package contractuc

import "github.com/naoyakurokawa/ddd_menta/core/domain/contractdm"

type ContractFindByIDUsecase interface {
	FindByID(contractID contractdm.ContractID) (*contractdm.Contract, error)
}

type ContractFindByIDUsecaseImpl struct {
	contractRepo contractdm.ContractRepository
}

func NewContractFindByIDUsecase(contractRepo contractdm.ContractRepository) ContractFindByIDUsecase {
	return &ContractFindByIDUsecaseImpl{contractRepo: contractRepo}
}

func (mu *ContractFindByIDUsecaseImpl) FindByID(contractID contractdm.ContractID) (*contractdm.Contract, error) {
	return mu.contractRepo.FindByID(contractID)
}
