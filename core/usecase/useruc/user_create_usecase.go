package useruc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

// UserUsecase user usecaseのinterface
type UserCreateUsecase interface {
	Create(name string, email string, password string, profile string) (*userdm.User, error)
}

type UserCreateUsecaseImpl struct {
	userRepo userdm.UserRepository
}

// user usecaseのコンストラクタ
func NewUserCreateUsecase(userRepo userdm.UserRepository) UserCreateUsecase {
	return &UserCreateUsecaseImpl{userRepo: userRepo}
}

// Create userを保存するときのユースケース
func (uu *UserCreateUsecaseImpl) Create(name string, email string, password string, profile string) (*userdm.User, error) {
	userID, err := userdm.NewUserID()
	if err != nil {
		return nil, err
	}
	emailIns, err := userdm.NewEmail(email)
	if err != nil {
		return nil, err
	}
	user, err := userdm.NewUser(userID, name, emailIns, password, profile)
	if err != nil {
		return nil, err
	}

	//最終的にinfraのCreateメソッドを実行することになる
	createdUser, err := uu.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
