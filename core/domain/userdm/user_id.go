package userdm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type UserID sharedvo.ID

func NewUserID() UserID {
	return UserID(sharedvo.NewID())
}

func NewUserIDByVal(strId string) (UserID, error) {
	id, err := sharedvo.NewIDByVal(strId)
	if err != nil {
		return UserID(""), xerrors.New("error NewUserIDByVal")
	}
	return UserID(id), nil
}

func (u UserID) Equals(u2 UserID) bool {
	return u.Value() == u2.Value()
}

func (u UserID) Value() string {
	return string(u)
}

func UserIDType(strUserID string) UserID {
	return UserID(strUserID)
}
