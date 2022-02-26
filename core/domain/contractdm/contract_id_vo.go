package contractdm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type ContractID sharedvo.ID

func NewContractID() ContractID {
	return ContractID(sharedvo.NewID())
}

func NewContractIDByVal(strId string) (ContractID, error) {
	id, err := sharedvo.NewIDByVal(strId)
	if err != nil {
		return ContractID(""), xerrors.New("error NewContractIDByVal")
	}
	return ContractID(id), nil
}

func NewEmptyContractID() ContractID {
	return ContractID(sharedvo.NewEmptyID())
}

func (i ContractID) Equals(i2 ContractID) bool {
	return i == i2
}

func (i ContractID) String() string {
	return string(i)
}
