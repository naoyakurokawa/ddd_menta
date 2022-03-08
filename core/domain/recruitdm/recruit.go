package recruitdm

import (
	"unicode/utf8"

	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type Recruit struct {
	recruitID     RecruitID
	userID        userdm.UserID
	title         string
	budget        uint16
	recruitType   RecruitType
	detail        string
	recruitStatus RecruitStatus
	createdAt     sharedvo.CreatedAt
	updatedAt     sharedvo.UpdatedAt
}

const (
	titleMaxLength  = 255
	detialMaxLength = 2000
	minBudget       = 1000
)

func NewRecruit(
	recruitID RecruitID,
	userID userdm.UserID,
	title string,
	budget uint16,
	recruitType RecruitType,
	detail string,
	recruitStatus RecruitStatus,
) (*Recruit, error) {
	if isEmpty(title) {
		return nil, xerrors.New("title must not be empty")
	}
	if isOver(title, titleMaxLength) {
		return nil, xerrors.Errorf("title must less than %d: %s", titleMaxLength, title)
	}
	if isLow(budget, minBudget) {
		return nil, xerrors.New("budget more than ¥1000")
	}
	if isEmpty(detail) {
		return nil, xerrors.New("detial must not be empty")
	}
	if isOver(detail, detialMaxLength) {
		return nil, xerrors.Errorf("title must less than %d: %s", titleMaxLength, title)
	}
	recruit := &Recruit{
		recruitID:     recruitID,
		userID:        userID,
		title:         title,
		budget:        budget,
		recruitType:   recruitType,
		detail:        detail,
		recruitStatus: recruitStatus,
		createdAt:     sharedvo.NewCreatedAt(),
		updatedAt:     sharedvo.NewUpdatedAt(),
	}

	return recruit, nil
}

func Reconstruct(
	recruitID string,
	userID string,
	title string,
	budget uint16,
	recruitType uint16,
	detail string,
	recruitStatus uint16,
) (*Recruit, error) {
	if isEmpty(title) {
		return nil, xerrors.New("title must not be empty")
	}
	if isOver(title, titleMaxLength) {
		return nil, xerrors.Errorf("title must less than %d: %s", titleMaxLength, title)
	}
	if isLow(budget, minBudget) {
		return nil, xerrors.New("budget more than ¥1000")
	}
	if isEmpty(detail) {
		return nil, xerrors.New("detial must not be empty")
	}
	if isOver(detail, detialMaxLength) {
		return nil, xerrors.Errorf("title must less than %d: %s", titleMaxLength, title)
	}
	castedRecruitID, err := NewRecruitIDByVal(recruitID)
	if err != nil {
		return nil, xerrors.New("error NewRecruitIDByVal")
	}
	castedUserID, err := userdm.NewUserIDByVal(userID)
	if err != nil {
		return nil, xerrors.New("error NewUserIDByVal")
	}
	recruitTypeIns, err := NewRecruitType(recruitType)
	if err != nil {
		return nil, xerrors.New("error NewRecruitType")
	}
	recruitStatusIns, err := NewRecruitStatus(recruitStatus)
	if err != nil {
		return nil, xerrors.New("error NewRecruitStatus")
	}
	recruit := &Recruit{
		recruitID:     castedRecruitID,
		userID:        castedUserID,
		title:         title,
		budget:        budget,
		recruitType:   recruitTypeIns,
		detail:        detail,
		recruitStatus: recruitStatusIns,
		createdAt:     sharedvo.NewCreatedAt(),
		updatedAt:     sharedvo.NewUpdatedAt(),
	}

	return recruit, nil
}

func (r *Recruit) RecruitID() RecruitID {
	return r.recruitID
}

func (r *Recruit) UserID() userdm.UserID {
	return r.userID
}

func (r *Recruit) Title() string {
	return r.title
}

func (r *Recruit) Budget() uint16 {
	return r.budget
}

func (r *Recruit) RecruitType() RecruitType {
	return r.recruitType
}

func (r *Recruit) Detail() string {
	return r.detail
}

func (r *Recruit) RecruitStatus() RecruitStatus {
	return r.recruitStatus
}

func isEmpty(s string) bool {
	return len(s) == 0
}

func isZero(u uint16) bool {
	return u == 0
}

func isOver(s string, maxlength int) bool {
	return utf8.RuneCountInString(s) > maxlength
}

func isLow(u, min uint16) bool {
	return min > u
}
