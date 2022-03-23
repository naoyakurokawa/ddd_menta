package userdm

type canRegisterEmailDomainService struct {
	userRepository UserRepository
}

func NewCanRegisterEmailDomainService(userRepository UserRepository) *canRegisterEmailDomainService {
	return &canRegisterEmailDomainService{
		userRepository: userRepository,
	}
}

func (s *canRegisterEmailDomainService) Exec(email Email) bool {
	user, err := s.userRepository.FetchByEmail(email)
	if err != nil && err.Error() != "record not found" {
		return false
	}
	if user == nil {
		return true
	}
	return false
}
