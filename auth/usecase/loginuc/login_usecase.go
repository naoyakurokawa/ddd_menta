package loginuc

import (
	"github.com/naoyakurokawa/ddd_menta/auth/domain/authuserdm"
	"github.com/naoyakurokawa/ddd_menta/auth/infrastructure/jwt"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/customerrors"
)

type LoginUsecase interface {
	Login(
		email string,
		password string,
	) (string, error)
}

type LoginUsecaseImpl struct {
	userRepo userdm.UserRepository
}

func NewLoginUsecase(userRepo userdm.UserRepository) LoginUsecase {
	return &LoginUsecaseImpl{userRepo: userRepo}
}

func (lu *LoginUsecaseImpl) Login(
	email string,
	password string,
) (string, error) {
	emailIns, err := userdm.NewEmail(email)
	if err != nil {
		return "", err
	}
	registeredUser, err := lu.userRepo.FetchByEmail(emailIns)
	if err != nil {
		return "", err
	}

	user, err := authuserdm.Reconstruct(
		email,
		password,
	)

	if err != nil {
		return "", err
	}

	if !user.VerifyPassword(registeredUser.Password().Value()) {
		return "", customerrors.NewUnauthorized()
	}

	t, err := jwt.NewToken(registeredUser.UserID().String(), registeredUser.Email().Value())
	if err != nil {
		return "", err
	}

	return t, nil
}
