package mentordm

type MentorRepository interface {
	Create(mentor *Mentor) (*Mentor, error)
	FindByID(mentorID MentorID) (*Mentor, error)
}
