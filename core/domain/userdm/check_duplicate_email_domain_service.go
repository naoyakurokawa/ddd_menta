package userdm

type checkDuplicateEmailDomainService struct {
	userRepository UserRepository
}

func NewCheckDuplicateEmailDomainService(userRepository UserRepository) *checkDuplicateEmailDomainService {
	return &checkDuplicateEmailDomainService{
		userRepository: userRepository,
	}
}

func (s *checkDuplicateEmailDomainService) Exec(email Email) bool {
	user, err := s.userRepository.FetchByEmail(email)
	if err != nil && user != nil {
		return false
	}
	return user == nil
}
