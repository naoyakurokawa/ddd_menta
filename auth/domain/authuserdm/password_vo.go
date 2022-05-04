package authuserdm

import (
	"unicode"

	"github.com/naoyakurokawa/ddd_menta/customerrors"
)

type Password string

const passwordMinLength = 12

func NewPassword(password string) (Password, error) {
	if len(password) == 0 {
		return Password(""), customerrors.NewInvalidParameter()
	}

	if len(password) < passwordMinLength {
		return Password(""), customerrors.NewInvalidParameter()
	}

	if !checkIncludeSpace(password) {
		return Password(""), customerrors.NewInvalidParameter()
	}

	if !checkIncludeAlphabet(password) {
		return Password(""), customerrors.NewInvalidParameter()
	}

	if !checkIncludeNumber(password) {
		return Password(""), customerrors.NewInvalidParameter()
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
