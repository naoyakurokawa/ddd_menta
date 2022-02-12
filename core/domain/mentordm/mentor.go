package mentordm

import (
	"unicode/utf8"

	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type Mentor struct {
	mentorID     MentorID
	userID       userdm.UserID
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
	mentorID MentorID,
	userID userdm.UserID,
	title string,
	mainImg string,
	subImg string,
	category string,
	detial string,
	mentorSkills []MentorSkill,
	plans []Plan,
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
		mentorID:     mentorID,
		userID:       userID,
		title:        title,
		mainImg:      mainImg,
		subImg:       subImg,
		category:     category,
		detial:       detial,
		createdAt:    sharedvo.NewCreatedAt(),
		mentorSkills: mentorSkills,
		plans:        plans,
	}

	return mentor, nil
}

func (m *Mentor) UserID() userdm.UserID {
	return m.userID
}

func (m *Mentor) MentorID() MentorID {
	return m.mentorID
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

func (m *Mentor) CreatedAt() sharedvo.CreatedAt {
	return m.createdAt
}

func (m *Mentor) MentorSkills() []MentorSkill {
	return m.mentorSkills
}

func (m *Mentor) Plans() []Plan {
	return m.plans
}

func (m *Mentor) AddMentorSkill(
	tag string,
	assessment uint16,
	experienceYears ExperienceYears,
) (*Mentor, error) {
	mentorSkillID := NewMentorSkillID()
	mentorSkill, err := NewMentorSkill(
		mentorSkillID,
		tag,
		assessment,
		experienceYears,
	)
	if err != nil {
		return nil, xerrors.New("error NewMentorSkill")
	}
	m.mentorSkills = append(m.mentorSkills, *mentorSkill)

	return m, nil
}

func (m *Mentor) AddPlan(
	title string,
	category string,
	tag string,
	detial string,
	planType PlanType,
	price uint16,
	planStatus PlanStatus,
) (*Mentor, error) {
	planID := NewPlanID()
	plan, err := NewPlan(
		planID,
		title,
		category,
		tag,
		detial,
		planType,
		price,
		planStatus,
	)
	if err != nil {
		return nil, xerrors.New("error newPlan")
	}
	m.plans = append(m.plans, *plan)

	return m, nil
}
