package mentordm

type checkNumMentorSkillDomainService struct {
	mentorRepository MentorRepository
}

func NewCheckNumMentorSkillDomainService(mentorRepository MentorRepository) *checkNumMentorSkillDomainService {
	return &checkNumMentorSkillDomainService{
		mentorRepository: mentorRepository,
	}
}

func (s *checkNumMentorSkillDomainService) Exec(requestMentorID MentorID) bool {
	mentor, err := s.mentorRepository.FindByID(requestMentorID)
	if err != nil {
		return false
	}
	if len(mentor.MentorSkills()) < 5 {
		return false
	}
	return true
}
