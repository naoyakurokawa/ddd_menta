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

	recruit, err := ru.recruitRepo.FetchByID(recruitIDIns)
	if err != nil {
		return err
	}

	if !recruit.IsPublished() {
		return xerrors.New("This recruit is not active")
	}

	mentor, err := ru.mentorRepo.FindByID(mentorIDIns)
	if err != nil {
		return err
	}

	if !mentor.CanSuggestion() {
		return xerrors.New("Can be suggested if you have 5 or more mentor skills")
	}

	return ru.suggestionRepo.Create(suggestion)
}
