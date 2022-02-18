package mentordm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type MentorID sharedvo.ID

func NewMentorID() MentorID {
	return MentorID(sharedvo.NewID())
}

func NewMentorIDByVal(strId string) (MentorID, error) {
	id, err := sharedvo.NewIDByVal(strId)
	if err != nil {
		return MentorID(""), xerrors.New("error NewMentorIDByVal")
	}
	return MentorID(id), nil
}

func NewEmptyMentorID() MentorID {
	return MentorID(sharedvo.NewEmptyID())
}

func (i MentorID) Equals(i2 MentorID) bool {
	return i == i2
}
