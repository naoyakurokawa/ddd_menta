package mentordm

import (
	"github.com/google/uuid"
)

type MentorSkillID string

func newMentorSkillID() (MentorSkillID, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return MentorSkillID(""), err
	}
	us := u.String()

	return MentorSkillID(us), nil
}

func (m MentorSkillID) Equals(m2 MentorSkillID) bool {
	return m.Value() == m2.Value()
}

func (m MentorSkillID) Value() string {
	return string(m)
}

func MentorSkillIDType(strMentorSkillID string) MentorSkillID {
	return MentorSkillID(strMentorSkillID)
}
