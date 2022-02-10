package mentordm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type MentorSkillID sharedvo.ID

func NewMentorSkillID() MentorSkillID {
	return MentorSkillID(sharedvo.NewID())
}

func NewMentorSkillIDByVal(srtID string) (MentorSkillID, error) {
	id, err := sharedvo.NewIDByVal(srtID)
	if err != nil {
		return MentorSkillID(""), xerrors.New("error NewMentorIDByVal")
	}
	return MentorSkillID(id), nil
}

func NewEmptyMentorSkillID() MentorSkillID {
	return MentorSkillID(sharedvo.NewEmptyID())
}
