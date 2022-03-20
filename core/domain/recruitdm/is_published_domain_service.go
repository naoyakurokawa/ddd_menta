package recruitdm

type isPublishedDomainServiceDomainService struct {
	recruitRepository RecruitRepository
}

func NewIsPublishedDomainServiceDomainService(recruitRepository RecruitRepository) *isPublishedDomainServiceDomainService {
	return &isPublishedDomainServiceDomainService{
		recruitRepository: recruitRepository,
	}
}

func (s *isPublishedDomainServiceDomainService) Exec(requestRecruitID RecruitID) bool {
	request, err := s.recruitRepository.FindByID(requestRecruitID)
	if err != nil {
		return false
	}
	if request.recruitStatus == Published {
		return true
	}
	return false
}
