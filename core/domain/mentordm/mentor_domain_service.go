package mentordm

type mentorDomainService struct {
	mentorRepository MentorRepository
}

func NewMentorDomainService(mentorRepository MentorRepository) *mentorDomainService {
	return &mentorDomainService{
		mentorRepository: mentorRepository,
	}
}

func (s *mentorDomainService) IsActivePlan(requestMentorID MentorID, requestPlanID PlanID) bool {
	mentor, err := s.mentorRepository.FindByID(requestMentorID)
	if err != nil {
		return false
	}
	for _, p := range mentor.plans {
		if p.planID == requestPlanID && p.planStatus == Active {
			return true
		}
	}
	return false
}
