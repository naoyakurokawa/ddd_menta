package mentordm

import (
	"unicode/utf8"

	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type MentorSkill struct {
	mentorSkillID   MentorSkillID
	tag             string
	assessment      uint16
	experienceYears ExperienceYears
	createdAt       sharedvo.CreatedAt
}

const (
	tagMaxLength     = 30
	assessmentMinNum = 1
	assessmentMaxNum = 5
)

func NewMentorSkill(
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

	mentorSkill := &MentorSkill{
		mentorSkillID:   mentorSkillID,
		tag:             tag,
		assessment:      assessment,
		experienceYears: experienceYears,
		createdAt:       sharedvo.NewCreatedAt(),
	}

	return mentorSkill, nil
}

func (m *MentorSkill) MentorSkillID() MentorSkillID {
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

func (u *MentorSkill) CreatedAt() sharedvo.CreatedAt {
	return u.createdAt
}
