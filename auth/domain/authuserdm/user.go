package authuserdm

import "golang.org/x/crypto/bcrypt"

type User struct {
	email    Email
	password Password
}

func NewUser(email Email, password Password) (*User, error) {
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
	emailIns, err := NewEmail(email)
	if err != nil {
		return nil, err
	}
	passwordIns, err := NewPassword(password)
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

func (u *User) Password() Password {
	return u.password
}
