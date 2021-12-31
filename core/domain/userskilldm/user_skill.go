package userskilldm

import (
	"time"

	"unicode/utf8"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type UserSkill struct {
	userSkillID     UserSkillID
	userID          userdm.UserID
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
func NewUserSkill(userSkillID UserSkillID, userID userdm.UserID, tag string, assessment uint16, experienceYears ExperienceYears) (*UserSkill, error) {
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

	userSkill := &UserSkill{
		userSkillID:     userSkillID,
		userID:          userID,
		tag:             tag,
		assessment:      assessment,
		experienceYears: experienceYears,
		createdAt:       now,
	}

	return userSkill, nil
}

func (u *UserSkill) UserSkillID() UserSkillID {
	return u.userSkillID
}

func (u *UserSkill) UserID() userdm.UserID {
	return u.userID
}

func (u *UserSkill) Tag() string {
	return u.tag
}

func (u *UserSkill) Assessment() uint16 {
	return u.assessment
}

func (u *UserSkill) ExperienceYears() ExperienceYears {
	return u.experienceYears
}

func (u *UserSkill) CreatedAt() time.Time {
	return u.createdAt
}
