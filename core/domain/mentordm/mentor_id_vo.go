package mentordm

import (
	"github.com/google/uuid"
)

type MentorID string

func NewMentorID() (MentorID, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return MentorID(""), err
	}
	us := u.String()

	return MentorID(us), nil
}

func (m MentorID) Equals(m2 MentorID) bool {
	return m.Value() == m2.Value()
}

func (m MentorID) Value() string {
	return string(m)
}

func MentorIDType(strMentorID string) MentorID {
	return MentorID(strMentorID)
}
