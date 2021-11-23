package useruc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

// UserUsecase user usecaseのinterface
type UserCreateUsecase interface {
	Create(name string, email string, password string, profile string, from []string, to []string, detail []string) (*userdm.User, error)
}

type UserCreateUsecaseImpl struct {
	userRepo userdm.UserRepository
}

// user usecaseのコンストラクタ
func NewUserCreateUsecase(userRepo userdm.UserRepository) UserCreateUsecase {
	return &UserCreateUsecaseImpl{userRepo: userRepo}
}

// Create userを保存するときのユースケース
func (uu *UserCreateUsecaseImpl) Create(name string, email string, password string, profile string, from []string, to []string, detail []string) (*userdm.User, error) {
	userID, err := userdm.NewUserID()
	if err != nil {
		return nil, err
	}
	emailIns, err := userdm.NewEmail(email)
	if err != nil {
		return nil, err
	}
	passwordIns, err := userdm.NewPassword(password)
	if err != nil {
		return nil, err
	}

	// userCareers := []userdm.UserCareer{}
	userCareers := make([]userdm.UserCareer, len(from))
	for i := 0; i < len(from); i++ {
		userCareerID, err := userdm.NewUserCareerID()
		if err != nil {
			return nil, err
		}
		userCareer, err := userdm.NewUserCareer(userCareerID, userID, from[i], to[i], detail[i])
		if err != nil {
			return nil, err
		}
		userCareers[i] = *userCareer
	}

	user, err := userdm.NewUser(userID, name, emailIns, passwordIns, profile, userCareers)
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
