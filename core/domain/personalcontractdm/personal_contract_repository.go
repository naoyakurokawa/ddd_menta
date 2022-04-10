package personalcontractdm

type PersonalContractRepository interface {
	Create(personalContract *PersonalContract) error
	FindByID(personalContractID PersonalContractID) (*PersonalContract, error)
}
