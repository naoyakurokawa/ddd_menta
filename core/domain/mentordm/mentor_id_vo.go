package mentordm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
)

type MentorID sharedvo.ID

func NewMentorID() MentorID {
	return MentorID(sharedvo.NewID())
}

func NewMentorIDByVal(id string) MentorID {
	return MentorID(id)
}

func NewEmptyMentorID() MentorID {
	return MentorID("")
}
