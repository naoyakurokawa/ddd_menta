package useruc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

type FetchUserByIdUsecase interface {
	FetchById(userID userdm.UserID) (*userdm.User, error)
}

type FetchUserByIdUsecaseImpl struct {
	userRepo userdm.UserRepository
}

// user usecaseのコンストラクタ
func NewUserFindByIDUsecase(userRepo userdm.UserRepository) FetchUserByIdUsecase {
	return &FetchUserByIdUsecaseImpl{userRepo: userRepo}
}

func (uu *FetchUserByIdUsecaseImpl) FetchById(userID userdm.UserID) (*userdm.User, error) {
	selectedUser, err := uu.userRepo.FetchById(userID)
	if err != nil {
		return nil, err
	}
	return selectedUser, nil
}
