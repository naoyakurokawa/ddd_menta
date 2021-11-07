package userdm

import (
	"github.com/google/uuid"
)

type UserCareerID string

func NewUserCareerID() (UserCareerID, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return UserCareerID(""), err
	}
	us := u.String()

	return UserCareerID(us), nil
}

func (u UserCareerID) Equals(u2 UserCareerID) bool {
	if u == u2 {
		return true
	} else {
		return false
	}
}

func (u UserCareerID) Value() string {
	return string(u)
}
