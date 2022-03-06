package contractdm

type ContractRepository interface {
	Create(contract *Contract) error
	FindByID(contractID ContractID) (*Contract, error)
	UpdateContractStatus(contractID *ContractID, contractStatus *ContractStatus) error
}
