package loginuc

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/naoyakurokawa/ddd_menta/config"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/crypto/bcrypt"
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

type jwtCustomClaims struct {
	UserID string `json:"userID"`
	Email  string `json:"email"`
	jwt.StandardClaims
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
	user, err := lu.userRepo.FetchByEmail(emailIns)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(password))
	if err != nil {
		fmt.Println("パスワードが一致しませんでした。：", err)
		return "", err
	}

	claims := &jwtCustomClaims{
		user.UserID().String(),
		user.Email().Value(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(config.SigningKey)
	if err != nil {
		return "", err
	}

	return t, nil
}
