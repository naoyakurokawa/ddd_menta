package useruc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

// UserUsecase user usecaseのinterface
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

// Create userを保存するときのユースケース
func (uu *UserUsecaseImpl) Create(name string, email string, password string, profile string) (*userdm.User, error) {
	userId, err := userdm.NewUserId()
	if err != nil {
		return nil, err
	}
	emailIns, err := userdm.NewEmail(email)
	if err != nil {
		return nil, err
	}
	user, err := userdm.NewUser(userId, name, emailIns, password, profile)
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
