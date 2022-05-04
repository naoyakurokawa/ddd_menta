package personalcontractuc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/personalcontractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
	"github.com/naoyakurokawa/ddd_menta/customerrors"
)

type CreatePersonalContractUsecase interface {
	Create(
		suggestionID string,
	) error
}

type CreatePersonalContractUsecaseImpl struct {
	personalContractRepo personalcontractdm.PersonalContractRepository
	suggestionRepo       suggestiondm.SuggestionRepository
}

func NewCreatePersonalContractUsecase(
	personalContractRepo personalcontractdm.PersonalContractRepository,
	suggestionRepo suggestiondm.SuggestionRepository,
) CreatePersonalContractUsecase {
	return &CreatePersonalContractUsecaseImpl{
		personalContractRepo: personalContractRepo,
		suggestionRepo:       suggestionRepo,
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

	// 提案のステータスが提案中でない場合は、契約不可
	suggestion, err := pu.suggestionRepo.FetchByID(suggestionIDIns)
	if err != nil {
		return err
	}
	if !suggestion.IsUnapproved() {
		return customerrors.NewInvalidParameter("suggestion must be unapproved")
	}

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
