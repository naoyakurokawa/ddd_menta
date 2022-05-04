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

	authUserEmailIns, err := authuserdm.NewEmail(email)
	if err != nil {
		return "", err
	}
	passwordIns, err := authuserdm.NewPassword(password)
	if err != nil {
		return "", err
	}

	user, err := authuserdm.NewUser(
		authUserEmailIns,
		passwordIns,
	)

	if err != nil {
		return "", err
	}

	if !user.VerifyPassword(registeredUser.Password().Value()) {
		return "", customerrors.NewUnauthorized("password is incorrect")
	}

	t, err := jwt.NewToken(registeredUser.UserID().String(), registeredUser.Email().Value())
	if err != nil {
		return "", err
	}

	return t, nil
}
