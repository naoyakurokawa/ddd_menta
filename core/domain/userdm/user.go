package userdm

import (
	"log"
	"reflect"
	"regexp"
	"time"

	"golang.org/x/xerrors"
)

type User struct {
	UserId    string
	Name      string
	Email     Email
	Password  string
	Profile   string
	CreatedAt time.Time
}

type Email string

var (
	emailFormat = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	emailRegExp = regexp.MustCompile(emailFormat)
)

const emailMaxLength = 256

// NewEmail emailのコンストラクタ
func NewEmail(email string) (Email, error) {
	if len(email) == 0 {
		return Email(""), xerrors.New("email must not be empty")
	}

	if len(email) > emailMaxLength {
		return Email(""), xerrors.Errorf("email must less than %d: %s", emailMaxLength, email)
	}

	if ok := emailRegExp.MatchString(email); !ok {
		return Email(""), xerrors.Errorf("invalid email format. email is %s", email)
	}

	return Email(email), nil
}

func (e Email) Value() string {
	return string(e)
}

func (e Email) Equals(e2 Email) bool {
	return reflect.DeepEqual(e, e2)
}

// NewUser userのコンストラクタ
func NewUser(name string, email string, password string, profile string) (*User, error) {
	//入力データチェック
	if len(name) == 0 {
		return nil, xerrors.New("name must not be empty")
	}
	if len(profile) == 0 {
		return nil, xerrors.New("profile must not be empty")
	}

	userId, err := NewUserId()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	now := time.Now()
	emailIns, err := NewEmail(email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user := &User{
		UserId:    userId.UserId,
		Name:      name,
		Email:     emailIns,
		Password:  password,
		Profile:   profile,
		CreatedAt: now,
	}

	return user, nil
}
