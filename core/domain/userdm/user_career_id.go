package userdm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
)

type UserCareerID sharedvo.ID

func NewUserCareerID() UserCareerID {
	return UserCareerID(sharedvo.NewID())
}

func (u UserCareerID) Equals(u2 UserCareerID) bool {
	return u.Value() == u2.Value()
}

func (u UserCareerID) Value() string {
	return string(u)
}
