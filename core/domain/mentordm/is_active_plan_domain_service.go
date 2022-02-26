package mentordm

type isActivePlanDomainService struct {
	mentorRepository MentorRepository
}

func NewIsActivePlanDomainService(mentorRepository MentorRepository) *isActivePlanDomainService {
	return &isActivePlanDomainService{
		mentorRepository: mentorRepository,
	}
}

func (s *isActivePlanDomainService) Exec(requestMentorID MentorID, requestPlanID PlanID) bool {
	mentor, err := s.mentorRepository.FindByID(requestMentorID)
	if err != nil {
		return false
	}
	for _, p := range mentor.plans {
		return p.isActive(requestPlanID)
	}
	return false
}
