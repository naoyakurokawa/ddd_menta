package userdm

import (
	"github.com/google/uuid"
)

type UserId string

func NewUserId() (UserId, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return UserId(""), err
	}
	us := u.String()

	return UserId(us), nil
}
