package suggestionuc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
)

type CreateSuggestionUsecase interface {
	Create(
		mentorID string,
		recruitID string,
		price uint32,
		suggestionType uint16,
		detail string,
		suggestionStatus uint16,
	) error
}

type CreateSuggestionUsecaseImpl struct {
	suggestionRepo suggestiondm.SuggestionRepository
}

func NewCreateSuggestionUsecase(
	suggestionRepo suggestiondm.SuggestionRepository,
) CreateSuggestionUsecase {
	return &CreateSuggestionUsecaseImpl{
		suggestionRepo: suggestionRepo,
	}
}

func (ru *CreateSuggestionUsecaseImpl) Create(
	mentorID string,
	recruitID string,
	price uint32,
	suggestionType uint16,
	detail string,
	suggestionStatus uint16,
) error {
	mentorIDIns, err := mentordm.NewMentorIDByVal(mentorID)
	if err != nil {
		return err
	}
	recruitIDIns, err := recruitdm.NewRecruitIDByVal(recruitID)
	if err != nil {
		return err
	}
	suggestionTypeIns, err := suggestiondm.NewSuggestionType(suggestionType)
	if err != nil {
		return err
	}
	suggestionStatusIns, err := suggestiondm.NewSuggestionStatus(suggestionStatus)
	if err != nil {
		return err
	}
	suggestionID := suggestiondm.NewSuggestionID()
	suggestion, err := suggestiondm.NewSuggestion(
		suggestionID,
		mentorIDIns,
		recruitIDIns,
		price,
		suggestionTypeIns,
		detail,
		suggestionStatusIns,
	)
	if err != nil {
		return err
	}

	//メンター募集のステータスが公開以外の場合は提案不可

	//メンターがスキル5つ未満の場合は提案不可

	return ru.suggestionRepo.Create(suggestion)
}
