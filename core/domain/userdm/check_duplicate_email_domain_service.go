package userdm

import (
	"errors"

	"github.com/naoyakurokawa/ddd_menta/customerrors"
)

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
	if err != nil && errors.Is(err, customerrors.Conflict) {
		return false
	}
	return user == nil
}
