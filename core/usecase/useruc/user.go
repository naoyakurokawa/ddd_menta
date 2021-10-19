package useruc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

// TaskUsecase task usecaseのinterface
type UserUsecase interface {
	Create(name string, email string, password string, profile string) (*userdm.User, error)
}

type UserUsecaseImpl struct {
	userRepo userdm.UserRepository
}

// user usecaseのコンストラクタ
func NewUserUsecase(userRepo userdm.UserRepository) UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}

// Create taskを保存するときのユースケース
func (uu *UserUsecaseImpl) Create(name string, email string, password string, profile string) (*userdm.User, error) {
	user, err := userdm.NewUser(name, email, password, profile)
	if err != nil {
		return nil, err
	}

	createdUser, err := uu.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
