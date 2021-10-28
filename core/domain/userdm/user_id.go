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
