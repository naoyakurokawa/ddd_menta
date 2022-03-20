package suggestionuc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
	"golang.org/x/xerrors"
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
	mentorRepo     mentordm.MentorRepository
	recruitRepo    recruitdm.RecruitRepository
}

func NewCreateSuggestionUsecase(
	suggestionRepo suggestiondm.SuggestionRepository,
	mentorRepo mentordm.MentorRepository,
	recruitRepo recruitdm.RecruitRepository,
) CreateSuggestionUsecase {
	return &CreateSuggestionUsecaseImpl{
		suggestionRepo: suggestionRepo,
		mentorRepo:     mentorRepo,
		recruitRepo:    recruitRepo,
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

	// メンター募集のステータスが公開以外の場合は提案不可
	isPublishedDomainServiceDomainService := recruitdm.NewIsPublishedDomainServiceDomainService(ru.recruitRepo)
	if !isPublishedDomainServiceDomainService.Exec(suggestion.RecruitID()) {
		return xerrors.New("This recruit is not active")
	}

	// メンターがスキル5つ未満の場合は提案不可
	checkNumMentorSkillDomainService := mentordm.NewCheckNumMentorSkillDomainService(ru.mentorRepo)
	if !checkNumMentorSkillDomainService.Exec(suggestion.MentorID()) {
		return xerrors.New("Can be suggested if you have 5 or more mentor skills")
	}

	return ru.suggestionRepo.Create(suggestion)
}
