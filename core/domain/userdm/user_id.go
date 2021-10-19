package userdm

import (
	"log"

	"github.com/google/uuid"
)

type UserId struct {
	UserId string
}

func NewUserId() (*UserId, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	us := u.String()
	user_id := &UserId{
		UserId: us,
	}

	return user_id, nil
}
