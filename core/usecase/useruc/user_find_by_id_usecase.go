package useruc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

// UserUsecase user usecaseのinterface
type UserFindByIDUsecase interface {
	FindByID(user *userdm.User) (*userdm.User, error)
}

type UserFindByIDUsecaseImpl struct {
	userRepo userdm.UserRepository
}

// user usecaseのコンストラクタ
func NewUserFindByIDUsecase(userRepo userdm.UserRepository) UserFindByIDUsecase {
	return &UserFindByIDUsecaseImpl{userRepo: userRepo}
}

func (uu *UserFindByIDUsecaseImpl) FindByID(user *userdm.User) (*userdm.User, error) {
	selectedUser, err := uu.userRepo.FindByID(user)
	if err != nil {
		return nil, err
	}
	return selectedUser, nil
}
