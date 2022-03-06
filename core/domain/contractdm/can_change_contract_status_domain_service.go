package contractdm

type canChangeContractStatusDomainService struct {
	contractRepository ContractRepository
}

func NewCanChangeContractStatusDomainService(
	contractRepository ContractRepository) *canChangeContractStatusDomainService {
	return &canChangeContractStatusDomainService{
		contractRepository: contractRepository,
	}
}

func (c *canChangeContractStatusDomainService) Exec(contractID ContractID, requestContractStatus ContractStatus) bool {
	contract, err := c.contractRepository.FindByID(contractID)
	if err != nil {
		return false
	}
	return requestContractStatus.Uint16() == contract.ContractStatus().Uint16()+1
}
