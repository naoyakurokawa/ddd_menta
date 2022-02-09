package mentordm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type MentorID sharedvo.ID

func NewMentorID() MentorID {
	return MentorID(sharedvo.NewID())
}

func NewMentorIDByVal(id string) (MentorID, error) {
	ID, err := sharedvo.NewIDByVal(id)
	if err != nil {
		return MentorID(""), xerrors.New("error NewMentorIDByVal")
	}
	return MentorID(ID), nil
}

func NewEmptyMentorID() MentorID {
	return MentorID(sharedvo.NewEmptyID())
}
