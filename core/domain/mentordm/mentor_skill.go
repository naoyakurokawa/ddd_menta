package mentordm

import (
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
	if isEmpty(tag) {
		return nil, xerrors.New("tag must not be empty")
	}
	if isOver(tag, tagMaxLength) {
		return nil, xerrors.Errorf("tag must less than %d: %s", tagMaxLength, tag)
	}
	if isBetween(assessment, assessmentMinNum, assessmentMaxNum) {
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

func ReconstructMentorSkill(
	mentorSkillID string,
	tag string,
	assessment uint16,
	experienceYears uint16,
) (*MentorSkill, error) {
	if isEmpty(tag) {
		return nil, xerrors.New("tag must not be empty")
	}
	if isOver(tag, tagMaxLength) {
		return nil, xerrors.Errorf("tag must less than %d: %s", tagMaxLength, tag)
	}
	if isBetween(assessment, assessmentMinNum, assessmentMaxNum) {
		return nil, xerrors.Errorf("assessment must between %d and %d", assessmentMinNum, assessmentMaxNum)
	}

	castedMentorSkillID, err := NewMentorSkillIDByVal(mentorSkillID)
	if err != nil {
		return nil, xerrors.New("error NewMentorSkillIDByVal")
	}
	experienceYearsIns, err := NewExperienceYears(experienceYears)
	if err != nil {
		return nil, xerrors.New("error NewExperienceYears")
	}
	mentorSkill := &MentorSkill{
		mentorSkillID:   castedMentorSkillID,
		tag:             tag,
		assessment:      assessment,
		experienceYears: experienceYearsIns,
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
