package mentordm

import (
	"time"
	"unicode/utf8"

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
	createdAt    time.Time
	plans        []Plan
	mentorSkills []MentorSkill
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
	plans []Plan,
	mentorSkills []MentorSkill,
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

	now := time.Now()

	mentor := &Mentor{
		userID:       userID,
		mentorID:     mentorID,
		title:        title,
		mainImg:      mainImg,
		subImg:       subImg,
		category:     category,
		detial:       detial,
		plans:        plans,
		mentorSkills: mentorSkills,
		createdAt:    now,
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

func (m *Mentor) NewPlan(
	title string,
	category string,
	tag string,
	detial string,
	planType uint16,
	price uint16,
	planStatus uint16,
) (*Plan, error) {
	planID, err := newPlanID()
	if err != nil {
		return nil, err
	}

	planTypeIns, err := newPlanType(planType)
	if err != nil {
		return nil, err
	}

	planStatusIns, err := newPlanStatus(planStatus)
	if err != nil {
		return nil, err
	}

	plan, err := newPlan(
		planID,
		title,
		category,
		tag,
		detial,
		planTypeIns,
		price,
		planStatusIns,
	)
	return plan, nil
}
