package personalcontractuc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/personalcontractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
)

type CreatePersonalContractUsecase interface {
	Create(
		suggestionID string,
	) error
}

type CreatePersonalContractUsecaseImpl struct {
	personalContractRepo personalcontractdm.PersonalContractRepository
}

func NewCreatePersonalContractUsecase(
	personalContractRepo personalcontractdm.PersonalContractRepository,
) CreatePersonalContractUsecase {
	return &CreatePersonalContractUsecaseImpl{
		personalContractRepo: personalContractRepo,
	}
}

func (pu *CreatePersonalContractUsecaseImpl) Create(
	suggestionID string,
) error {
	personalContractID := personalcontractdm.NewPersonalContractID()
	suggestionIDIns, err := suggestiondm.NewSuggestionIDByVal(suggestionID)
	if err != nil {
		return err
	}
	// castedPersonalContractStatus, err := personalcontractdm.NewPersonalContractStatus(personalContractStatus)
	// if err != nil {
	// 	return err
	// }

	// メンティーが提案を承認した際に生成されるためStatusは契約中
	personalContract, err := personalcontractdm.GenWhenCreate(
		personalContractID,
		suggestionIDIns,
	)
	if err != nil {
		return err
	}
	return pu.personalContractRepo.Create(personalContract)
}
