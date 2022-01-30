package mentordm

import (
	"unicode/utf8"

	"github.com/naoyakurokawa/ddd_menta/core/domain/shared"
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
	createdAt    shared.CreatedAt
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
	tag []string,
	assessment []uint16,
	experienceYears []uint16,
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

	// メンタースキル
	mentorSkills := make([]MentorSkill, len(tag))
	for i := 0; i < len(tag); i++ {
		mentorSkillID, err := newMentorSkillID()
		if err != nil {
			return nil, xerrors.New("error newMentorSkillID")
		}
		experienceYearsIns, err := newExperienceYears(experienceYears[i])
		if err != nil {
			return nil, xerrors.New("error newExperienceYears")
		}
		mentorSkill, err := newMentorSkill(
			mentorSkillID,
			tag[i],
			assessment[i],
			experienceYearsIns,
		)
		if err != nil {
			return nil, xerrors.New("error newMentorSkill")
		}
		mentorSkills[i] = *mentorSkill
	}

	mentor := &Mentor{
		userID:       userID,
		mentorID:     mentorID,
		title:        title,
		mainImg:      mainImg,
		subImg:       subImg,
		category:     category,
		detial:       detial,
		createdAt:    shared.GetCurrentTime(),
		mentorSkills: mentorSkills,
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
		planID, err := newPlanID()
		if err != nil {
			return nil, xerrors.New("error newPlanID")
		}
		planType, err := newPlanType(planType[i])
		if err != nil {
			return nil, xerrors.New("error newPlanType")
		}
		planStatus, err := newPlanStatus(planStatus[i])
		if err != nil {
			return nil, xerrors.New("error newPlanStatus")
		}
		plan, err := newPlan(
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
	mentor := &Mentor{plans: plans}
	return mentor, nil
}
