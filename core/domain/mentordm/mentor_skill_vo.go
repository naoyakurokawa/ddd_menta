package mentordm

import (
	"strconv"
	"time"
	"unicode/utf8"

	"golang.org/x/xerrors"
)

type MentorSkill struct {
	mentorID        MentorID
	mentorSkillID   MentorSkillID
	tag             string
	assessment      uint16
	experienceYears ExperienceYears
	createdAt       time.Time
}

const (
	tagMaxLength     = 30
	assessmentMinNum = 1
	assessmentMaxNum = 5
)

// NewUserSkill userSkillのコンストラクタ
func NewUserSkill(
	mentorID MentorID,
	mentorSkillID MentorSkillID,
	tag string,
	assessment uint16,
	experienceYears ExperienceYears,
) (*MentorSkill, error) {
	//入力データチェック
	if len(tag) == 0 {
		return nil, xerrors.New("tag must not be empty")
	}
	if utf8.RuneCountInString(tag) > tagMaxLength {
		return nil, xerrors.Errorf("tag must less than %d: %s", tagMaxLength, tag)
	}
	if assessment < assessmentMinNum || assessmentMaxNum < assessment {
		return nil, xerrors.Errorf("assessment must between %d and %d", assessmentMinNum, assessmentMaxNum)
	}

	now := time.Now()

	mentorSkill := &MentorSkill{
		mentorID:        mentorID,
		mentorSkillID:   mentorSkillID,
		tag:             tag,
		assessment:      assessment,
		experienceYears: experienceYears,
		createdAt:       now,
	}

	return mentorSkill, nil
}

func (m *MentorSkill) MentorID() MentorID {
	return m.mentorID
}

func (m *MentorSkill) UserSkillID() MentorSkillID {
	return m.mentorSkillID
}

func (u *MentorSkill) Tag() string {
	return u.tag
}

func (u *MentorSkill) Assessment() uint16 {
	return u.assessment
}

func (u *MentorSkill) ExperienceYears() ExperienceYears {
	return u.experienceYears
}

func (u *MentorSkill) CreatedAt() time.Time {
	return u.createdAt
}

func (u *MentorSkill) AssessmentCastUint(assessment string) (uint16, error) {
	uintAssessment, err := strconv.ParseUint(assessment, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(uintAssessment), nil
}
