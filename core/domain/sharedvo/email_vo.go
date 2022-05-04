package sharedvo

import (
	"fmt"
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
		return Email(""), customerrors.NewInvalidParameter("email must not be empty")
	}

	if utf8.RuneCountInString(email) > emailMaxLength {
		return Email(""), customerrors.NewInvalidParameter(fmt.Sprintf("email must less than %d: %s", emailMaxLength, email))
	}

	if ok := emailRegExp.MatchString(email); !ok {
		return Email(""), customerrors.NewInvalidParameter(fmt.Sprintf("invalid email format. email is %s", email))
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
