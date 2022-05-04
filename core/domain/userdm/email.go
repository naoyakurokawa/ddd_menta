package userdm

import (
	"regexp"
	"unicode/utf8"

	"github.com/naoyakurokawa/ddd_menta/customerrors"
)

type Email string

var (
	emailFormat = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	emailRegExp = regexp.MustCompile(emailFormat)
)

const emailMaxLength = 255

// NewEmail emailのコンストラクタ
func NewEmail(email string) (Email, error) {
	if len(email) == 0 {
		return Email(""), customerrors.NewInvalidParameter()
	}

	if utf8.RuneCountInString(email) > emailMaxLength {
		return Email(""), customerrors.NewInvalidParameter()
	}

	if ok := emailRegExp.MatchString(email); !ok {
		return Email(""), customerrors.NewInvalidParameter()
	}

	return Email(email), nil
}

func (e Email) Value() string {
	return string(e)
}

func EmailType(strEmail string) Email {
	return Email(strEmail)
}

func (e Email) Equals(e2 Email) bool {
	return e.Value() == e2.Value()
}
