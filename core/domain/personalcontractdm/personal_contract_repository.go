package personalcontractdm

type PersonalContractRepository interface {
	Create(personalContract *PersonalContract) error
	FetchByID(personalContractID PersonalContractID) (*PersonalContract, error)
}
