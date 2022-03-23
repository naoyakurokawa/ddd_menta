package userdm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type UserSkillID sharedvo.ID

func NewUserSkillID() UserSkillID {
	return UserSkillID(sharedvo.NewID())
}

func NewUserSkillIDByVal(strId string) (UserSkillID, error) {
	id, err := sharedvo.NewIDByVal(strId)
	if err != nil {
		return UserSkillID(""), xerrors.New("error NewUserSkillIDByVal")
	}
	return UserSkillID(id), nil
}

func (u UserSkillID) Equals(u2 UserSkillID) bool {
	return u.Value() == u2.Value()
}

func (u UserSkillID) Value() string {
	return string(u)
}

func UserSkillIDType(strUserSkillID string) (UserSkillID, error) {
	if len(strUserSkillID) == 0 {
		return UserSkillID(""), xerrors.New("tag must not be empty")
	}
	return UserSkillID(strUserSkillID), nil
}
