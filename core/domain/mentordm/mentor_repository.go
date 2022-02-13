package mentordm

type MentorRepository interface {
	Create(mentor *Mentor) error
	FindByID(mentorID MentorID) (*Mentor, error)
}
