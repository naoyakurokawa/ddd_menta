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

func (m *Mentor) MentorSkills() []MentorSkill {
	return m.mentorSkills
}

func (m *Mentor) Plans() []Plan {
	return m.plans
}

func (m *Mentor) AddMentorSkill(
	tag []string,
	assessment []uint16,
	experienceYears []uint16,
) (*Mentor, error) {
	mentorSkills := make([]MentorSkill, len(tag))
	for i := 0; i < len(tag); i++ {
		mentorSkillID := NewMentorSkillID()
		experienceYears, err := NewExperienceYears(experienceYears[i])
		if err != nil {
			return nil, xerrors.New("error NewExperienceYears")
		}
		mentorSkill, err := NewMentorSkill(
			mentorSkillID,
			tag[i],
			assessment[i],
			experienceYears,
		)
		if err != nil {
			return nil, xerrors.New("error NewMentorSkill")
		}
		mentorSkills[i] = *mentorSkill
	}
	m.mentorSkills = mentorSkills
	return m, nil
}

func (m *Mentor) AddPlan(
	title []string,
	category []string,
	tag []string,
	detial []string,
	planType []uint16,
	price []uint16,
	planStatus []uint16,
) (*Mentor, error) {
	plans := make([]Plan, len(title))
	for i := 0; i < len(title); i++ {
		planID := NewPlanID()
		planType, err := NewPlanType(planType[i])
		if err != nil {
			return nil, xerrors.New("error newPlanType")
		}
		planStatus, err := NewPlanStatus(planStatus[i])
		if err != nil {
			return nil, xerrors.New("error newPlanStatus")
		}
		plan, err := NewPlan(
			planID,
			title[i],
			category[i],
			tag[i],
			detial[i],
			planType,
			price[i],
			planStatus,
		)
		if err != nil {
			return nil, xerrors.New("error newPlan")
		}
		plans[i] = *plan
	}
	m.plans = plans
	return m, nil
}
