package personalcontractdm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
)

func GenWhenCreate(
	personalContractID PersonalContractID,
	suggestionID suggestiondm.SuggestionID,
) (*PersonalContract, error) {
	return NewPersonalContract(
		personalContractID,
		suggestionID,
		UnderContract,
	)
}
