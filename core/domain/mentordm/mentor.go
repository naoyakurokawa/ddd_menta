package mentordm

import (
	"unicode/utf8"

	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type Mentor struct {
	userID       userdm.UserID
	mentorID     MentorID
	title        string
	mainImg      string
	subImg       string
	category     string
	detial       string
	createdAt    sharedvo.CreatedAt
	mentorSkills []MentorSkill
	plans        []Plan
}

const titleMaxLength = 255
const detialMaxLength = 2000

func NewMentor(
	userID userdm.UserID,
	mentorID MentorID,
	title string,
	mainImg string,
	subImg string,
	category string,
	detial string,
) (*Mentor, error) {
	//入力データチェック
	if utf8.RuneCountInString(title) > titleMaxLength {
		return nil, xerrors.Errorf("title must less than %d: %s", titleMaxLength, title)
	}
	if len(title) == 0 {
		return nil, xerrors.New("title must not be empty")
	}
	if len(detial) == 0 {
		return nil, xerrors.New("detial must not be empty")
	}
	if utf8.RuneCountInString(detial) > detialMaxLength {
		return nil, xerrors.Errorf("detial must less than %d: %s", detialMaxLength, detial)
	}

	mentor := &Mentor{
		userID:    userID,
		mentorID:  mentorID,
		title:     title,
		mainImg:   mainImg,
		subImg:    subImg,
		category:  category,
		detial:    detial,
		createdAt: sharedvo.NewCreatedAt(),
	}

	return mentor, nil
}

func (m *Mentor) UserID() userdm.UserID {
	return m.userID
}

func (m *Mentor) Title() string {
	return m.title
}

func (m *Mentor) MainImg() string {
	return m.mainImg
}

func (m *Mentor) SubImg() string {
	return m.subImg
}

func (m *Mentor) Category() string {
	return m.category
}

func (m *Mentor) Detail() string {
	return m.detial
}
