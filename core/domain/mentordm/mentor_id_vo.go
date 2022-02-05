package mentordm

import (
	"github.com/google/uuid"
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
)

type MentorID sharedvo.ID

func NewMentorID() MentorID {
	return MentorID(uuid.New().String())
}

func NewMentorIDByVal(id string) MentorID {
	return MentorID(id)
}

func NewEmptyMentorID() MentorID {
	return MentorID("")
}
