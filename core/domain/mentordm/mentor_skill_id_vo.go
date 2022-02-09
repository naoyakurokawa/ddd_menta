package mentordm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type MentorSkillID sharedvo.ID

func NewMentorSkillID() MentorSkillID {
	return MentorSkillID(sharedvo.NewID())
}

func NewMentorSkillIDByVal(id string) (MentorSkillID, error) {
	ID, err := sharedvo.NewIDByVal(id)
	if err != nil {
		return MentorSkillID(""), xerrors.New("error NewMentorIDByVal")
	}
	return MentorSkillID(ID), nil
}

func NewEmptyMentorSkillID() MentorSkillID {
	return MentorSkillID(sharedvo.NewEmptyID())
}
