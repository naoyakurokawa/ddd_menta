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

func IsEqualUserID(userID1, userID2 UserID) bool {
	if userID1 == userID2 {
		return true
	} else {
		return false
	}
}
