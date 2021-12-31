package userskilldm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type UserSkillID string

func NewUserSkillID() (UserSkillID, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return UserSkillID(""), err
	}
	us := u.String()

	return UserSkillID(us), nil
}

func (u UserSkillID) Equals(u2 UserSkillID) bool {
	return u.Value() == u2.Value()
}

func (u UserSkillID) Value() string {
	return string(u)
}

func UserIDType(strUserSkillID string) (UserSkillID, error) {
	if len(strUserSkillID) == 0 {
		return UserSkillID(""), xerrors.New("tag must not be empty")
	}
	return UserSkillID(strUserSkillID), nil
}
