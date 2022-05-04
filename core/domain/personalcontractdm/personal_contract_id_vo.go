package personalcontractdm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"github.com/naoyakurokawa/ddd_menta/customerrors"
)

type PersonalContractID sharedvo.ID

func NewPersonalContractID() PersonalContractID {
	return PersonalContractID(sharedvo.NewID())
}

func NewPersonalContractIDByVal(strId string) (PersonalContractID, error) {
	id, err := sharedvo.NewIDByVal(strId)
	if err != nil {
		return PersonalContractID(""), customerrors.NewInvalidParameter("error NewPersonalContractIDByVal")
	}
	return PersonalContractID(id), nil
}

func NewEmptyPersonalContractID() PersonalContractID {
	return PersonalContractID(sharedvo.NewEmptyID())
}

func (i PersonalContractID) Equals(i2 PersonalContractID) bool {
	return i == i2
}

func (i PersonalContractID) String() string {
	return string(i)
}
