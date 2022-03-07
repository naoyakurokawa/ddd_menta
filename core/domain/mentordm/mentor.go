package mentordm

import (
	"strconv"
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
	detail string,
	mentorSkills []MentorSkill,
	plans []Plan,
) (*Mentor, error) {
	//入力データチェック
	if isEmpty(title) {
		return nil, xerrors.New("title must not be empty")
	}
	if isOver(title, titleMaxLength) {
		return nil, xerrors.Errorf("title must less than %d: %s", titleMaxLength, title)
	}
	if isEmpty(detail) {
		return nil, xerrors.New("detial must not be empty")
	}
	if isOver(detail, detialMaxLength) {
		return nil, xerrors.Errorf("title must less than %d: %s", titleMaxLength, title)
	}

	mentor := &Mentor{
		mentorID:     mentorID,
		userID:       userID,
		title:        title,
		mainImg:      mainImg,
		subImg:       subImg,
		category:     category,
		detial:       detail,
		createdAt:    sharedvo.NewCreatedAt(),
		mentorSkills: mentorSkills,
		plans:        plans,
	}

	return mentor, nil
}

func Reconstruct(
	mentorID string,
	userID string,
	title string,
	mainImg string,
	subImg string,
	category string,
	detial string,
	mentorSkills []MentorSkill,
	plans []Plan,
) (*Mentor, error) {
	//入力データチェック
	if isEmpty(title) {
		return nil, xerrors.New("title must not be empty")
	}
	if isOver(title, titleMaxLength) {
		return nil, xerrors.Errorf("title must less than %d: %s", titleMaxLength, title)
	}
	if isEmpty(detial) {
		return nil, xerrors.New("detial must not be empty")
	}
	if isOver(detial, detialMaxLength) {
		return nil, xerrors.Errorf("title must less than %d: %s", titleMaxLength, title)
	}

	castedMentorID, err := NewMentorIDByVal(mentorID)
	if err != nil {
		return nil, xerrors.New("error NewMentorIDByVal")
	}
	castedUserID, err := userdm.NewUserIDByVal(userID)
	if err != nil {
		return nil, xerrors.New("error NewUserIDByVal")
	}

	mentor := &Mentor{
		mentorID:     castedMentorID,
		userID:       castedUserID,
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

func MentorIDType(strMentorID string) MentorID {
	return MentorID(strMentorID)
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

func isEmpty(s string) bool {
	return len(s) == 0
}

func isZero(u uint16) bool {
	return u == 0
}

func isOver(s string, maxlength int) bool {
	return utf8.RuneCountInString(s) > maxlength
}

func isBetween(u uint16, minlength int, maxlength int) bool {
	return u < uint16(minlength) || uint16(maxlength) < u
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

func StrCastUint(str string) (uint16, error) {
	ui, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(ui), nil
}
