package suggestiondm

import (
	"unicode/utf8"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type Suggestion struct {
	suggestionID     SuggestionID
	mentorID         mentordm.MentorID
	recruitID        recruitdm.RecruitID
	price            uint32
	suggestionType   SuggestionType
	detail           string
	suggestionStatus SuggestionStatus
	createdAt        sharedvo.CreatedAt
	updatedAt        sharedvo.UpdatedAt
}

const (
	detailMaxLength = 2000
	minPrice        = 1000
)

func NewSuggestion(
	suggestionID SuggestionID,
	mentorID mentordm.MentorID,
	recruitID recruitdm.RecruitID,
	price uint32,
	suggestionType SuggestionType,
	detail string,
	suggestionStatus SuggestionStatus,
) (*Suggestion, error) {
	if isLow(price, minPrice) {
		return nil, xerrors.New("price more than ¥1000")
	}
	if isEmpty(detail) {
		return nil, xerrors.New("detail must not be empty")
	}
	if isOver(detail, detailMaxLength) {
		return nil, xerrors.Errorf("detail must less than %d: %s", detailMaxLength, detail)
	}
	suggestion := &Suggestion{
		suggestionID:     suggestionID,
		mentorID:         mentorID,
		recruitID:        recruitID,
		price:            price,
		suggestionType:   suggestionType,
		detail:           detail,
		suggestionStatus: suggestionStatus,
		createdAt:        sharedvo.NewCreatedAt(),
		updatedAt:        sharedvo.NewUpdatedAt(),
	}

	return suggestion, nil
}

func Reconstruct(
	suggestionID string,
	mentorID string,
	recruitID string,
	price uint32,
	suggestionType uint16,
	detail string,
	suggestionStatus uint16,
) (*Suggestion, error) {
	if isLow(price, minPrice) {
		return nil, xerrors.New("price more than ¥1000")
	}
	if isEmpty(detail) {
		return nil, xerrors.New("detail must not be empty")
	}
	if isOver(detail, detailMaxLength) {
		return nil, xerrors.Errorf("detail must less than %d: %s", detailMaxLength, detail)
	}
	castedSuggestionID, err := NewSuggestionIDByVal(suggestionID)
	if err != nil {
		return nil, xerrors.New("error NewRecruitIDByVal")
	}
	castedMentorID, err := mentordm.NewMentorIDByVal(mentorID)
	if err != nil {
		return nil, xerrors.New("error NewMentorIDByVal")
	}
	castedRecruitID, err := recruitdm.NewRecruitIDByVal(recruitID)
	if err != nil {
		return nil, xerrors.New("error NewRecruitIDByVal")
	}
	suggestionTypeIns, err := NewSuggestionType(suggestionType)
	if err != nil {
		return nil, xerrors.New("error NewSuggestionType")
	}
	suggestionStatusIns, err := NewSuggestionStatus(suggestionStatus)
	if err != nil {
		return nil, xerrors.New("error NewSuggestionStatus")
	}
	suggestion := &Suggestion{
		suggestionID:     castedSuggestionID,
		mentorID:         castedMentorID,
		recruitID:        castedRecruitID,
		price:            price,
		suggestionType:   suggestionTypeIns,
		detail:           detail,
		suggestionStatus: suggestionStatusIns,
		createdAt:        sharedvo.NewCreatedAt(),
		updatedAt:        sharedvo.NewUpdatedAt(),
	}

	return suggestion, nil
}

func (s *Suggestion) SuggestionID() SuggestionID {
	return s.suggestionID
}

func (s *Suggestion) MentorID() mentordm.MentorID {
	return s.mentorID
}

func (s *Suggestion) RecruitID() recruitdm.RecruitID {
	return s.recruitID
}

func (s *Suggestion) Price() uint32 {
	return s.price
}

func (s *Suggestion) SuggestionType() SuggestionType {
	return s.suggestionType
}

func (s *Suggestion) Detail() string {
	return s.detail
}

func (s *Suggestion) SuggestionStatus() SuggestionStatus {
	return s.suggestionStatus
}

func isEmpty(s string) bool {
	return len(s) == 0
}

func isOver(s string, maxlength int) bool {
	return utf8.RuneCountInString(s) > maxlength
}

func isLow(u, min uint32) bool {
	return min > u
}

func (s *Suggestion) IsUnapproved() bool {
	return s.suggestionStatus == Unapproved
}
