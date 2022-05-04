package authuserdm

import (
	"fmt"
	"unicode"

	"github.com/naoyakurokawa/ddd_menta/customerrors"
)

type Password string

const passwordMinLength = 12

func NewPassword(password string) (Password, error) {
	if len(password) == 0 {
		return Password(""), customerrors.NewInvalidParameter("password must not be empty")
	}

	if len(password) < passwordMinLength {
		return Password(""), customerrors.NewInvalidParameter(fmt.Sprintf("password must more than %d length", passwordMinLength))
	}

	if !checkIncludeSpace(password) {
		return Password(""), customerrors.NewInvalidParameter("password not use spase")
	}

	if !checkIncludeAlphabet(password) {
		return Password(""), customerrors.NewInvalidParameter("password must use alphabet more than one")
	}

	if !checkIncludeNumber(password) {
		return Password(""), customerrors.NewInvalidParameter("password must use number more than one")
	}

	return Password(password), nil
}

func checkIncludeAlphabet(password string) bool {
	for _, s := range password {
		if unicode.IsLower(s) || unicode.IsUpper(s) {
			return true
		}
	}
	return false
}

func checkIncludeNumber(password string) bool {
	for _, s := range password {
		if unicode.IsNumber(s) {
			return true
		}
	}
	return false
}

func checkIncludeSpace(password string) bool {
	for _, s := range password {
		if unicode.IsSpace(s) {
			return false
		}
	}
	return true
}

func (p Password) Value() string {
	return string(p)
}

func PasswordType(strPassword string) Password {
	return Password(strPassword)
}

func (p Password) Equals(p2 Password) bool {
	return p.Value() == p2.Value()
}
