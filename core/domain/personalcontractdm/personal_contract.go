package personalcontractdm

import (
	"strconv"

	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
	"github.com/naoyakurokawa/ddd_menta/customerrors"
)

type PersonalContract struct {
	personalContractID     PersonalContractID
	suggestionID           suggestiondm.SuggestionID
	personalContractStatus PersonalContractStatus
	createdAt              sharedvo.CreatedAt
	updatedAt              sharedvo.UpdatedAt
}

func NewPersonalContract(
	personalContractID PersonalContractID,
	suggestionID suggestiondm.SuggestionID,
	personalContractStatus PersonalContractStatus,
) (*PersonalContract, error) {

	personaContract := &PersonalContract{
		personalContractID:     personalContractID,
		suggestionID:           suggestionID,
		personalContractStatus: personalContractStatus,
		createdAt:              sharedvo.NewCreatedAt(),
		updatedAt:              sharedvo.NewUpdatedAt(),
	}

	return personaContract, nil
}

func Reconstruct(
	personalContractID string,
	suggestionID string,
	personalContractStatus uint16,
) (*PersonalContract, error) {
	castedPersonalContractID, err := NewPersonalContractIDByVal(personalContractID)
	if err != nil {
		return nil, err
	}
	castedSuggestionID, err := suggestiondm.NewSuggestionIDByVal(suggestionID)
	if err != nil {
		return nil, err
	}
	personalContractStatusIns, err := NewPersonalContractStatus(personalContractStatus)
	if err != nil {
		return nil, customerrors.NewInvalidParameter("PersonalContractStatus must be 1 or 2 or 3")
	}

	personaContract := &PersonalContract{
		personalContractID:     castedPersonalContractID,
		suggestionID:           castedSuggestionID,
		personalContractStatus: personalContractStatusIns,
		createdAt:              sharedvo.NewCreatedAt(),
		updatedAt:              sharedvo.NewUpdatedAt(),
	}

	return personaContract, nil
}

func (p *PersonalContract) PersonalContractID() PersonalContractID {
	return p.personalContractID
}

func (p *PersonalContract) SuggestionID() suggestiondm.SuggestionID {
	return p.suggestionID
}

func (p *PersonalContract) PersonalContractStatus() PersonalContractStatus {
	return p.personalContractStatus
}

func (p *PersonalContract) CreatedAt() sharedvo.CreatedAt {
	return p.createdAt
}

func (p *PersonalContract) UpdatedAt() sharedvo.UpdatedAt {
	return p.updatedAt
}

func StrCastUint(str string) (uint16, error) {
	ui, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(ui), nil
}
