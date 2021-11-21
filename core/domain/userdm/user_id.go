package userdm

import (
	"github.com/google/uuid"
)

type UserID string

func NewUserID() (UserID, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return UserID(""), err
	}
	us := u.String()

	return UserID(us), nil
}

func (u UserID) Equals(u2 UserID) bool {
	if u == u2 {
		return true
	} else {
		return false
	}
}

func (u UserID) Value() string {
	return string(u)
}

func UserIDType(strUserID string) UserID {
	return UserID(strUserID)
}
