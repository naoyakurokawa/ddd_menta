package useruc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

type FetchUserByEmailUsecase interface {
	FetchByEmail(email userdm.Email) (*userdm.User, error)
}

type FetchUserByEmailUsecaseImpl struct {
	userRepo userdm.UserRepository
}

// user usecaseのコンストラクタ
func NewFetchUserByEmailUsecase(userRepo userdm.UserRepository) FetchUserByEmailUsecase {
	return &FetchUserByEmailUsecaseImpl{userRepo: userRepo}
}

func (uu *FetchUserByEmailUsecaseImpl) FetchByEmail(email userdm.Email) (*userdm.User, error) {
	return uu.userRepo.FetchByEmail(email)
}
