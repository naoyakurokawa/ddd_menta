package authuserdm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	email    sharedvo.Email
	password sharedvo.Password
}

func NewUser(email sharedvo.Email, password sharedvo.Password) (*User, error) {
	user := &User{
		email:    email,
		password: password,
	}

	return user, nil
}

func Reconstruct(
	email string,
	password string,
) (*User, error) {
	emailIns, err := sharedvo.NewEmail(email)
	if err != nil {
		return nil, err
	}
	passwordIns, err := sharedvo.NewPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		email:    emailIns,
		password: passwordIns,
	}

	return user, nil
}

func (u *User) VerifyPassword(hashPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(u.Password()))
	if err != nil {
		return false
	}
	return true
}

func (u *User) Password() sharedvo.Password {
	return u.password
}
