package mentordm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
)

type MentorSkillID sharedvo.ID

func NewMentorSkillID() MentorSkillID {
	return MentorSkillID(sharedvo.NewID())
}

func NewMentorSkillIDByVal(id string) MentorSkillID {
	return MentorSkillID(id)
}

func NewEmptyMentorSkillID() MentorSkillID {
	return MentorSkillID("")
}
