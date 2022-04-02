package personalcontractdm

import (
	"strconv"

	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
	"golang.org/x/xerrors"
)

type PersonalContract struct {
	personalContractID     PersonalContractID
	suggestionID           suggestiondm.SuggestionID
	personalContractStatus PersonalContractStatus
	createdAt              sharedvo.CreatedAt
	updatedAt              sharedvo.UpdatedAt
}

func NewContract(
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
		return nil, xerrors.New("error NewPersonalContractIDByVal")
	}
	castedSuggestionID, err := suggestiondm.NewSuggestionIDByVal(suggestionID)
	if err != nil {
		return nil, xerrors.New("error NewSuggestionIDByVal")
	}
	personalContractStatusIns, err := NewPersonalContractStatus(personalContractStatus)
	if err != nil {
		return nil, xerrors.New("error NewPersonalContractStatus")
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

func (c *PersonalContract) PersonalContractID() PersonalContractID {
	return c.personalContractID
}

func (c *PersonalContract) SuggestionID() suggestiondm.SuggestionID {
	return c.suggestionID
}

func (c *PersonalContract) PersonalContractStatus() PersonalContractStatus {
	return c.personalContractStatus
}

func (c *PersonalContract) CreatedAt() sharedvo.CreatedAt {
	return c.createdAt
}

func (c *PersonalContract) UpdatedAt() sharedvo.UpdatedAt {
	return c.updatedAt
}

func StrCastUint(str string) (uint16, error) {
	ui, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(ui), nil
}

func (c *PersonalContract) CanUpdatePersonalContractStatus(
	requestPersonalContractStatus PersonalContractStatus,
) bool {
	return requestPersonalContractStatus.Uint16() == c.personalContractStatus.Uint16()+1
}
