package userdm

import (
	"time"

	"golang.org/x/xerrors"
)

type User struct {
	UserId    UserId
	Name      string
	Email     Email
	Password  string
	Profile   string
	CreatedAt time.Time
}

const nameMaxLength = 255
const profileMaxLength = 2000

// NewUser userのコンストラクタ
func NewUser(name string, email string, password string, profile string) (*User, error) {
	//入力データチェック
	if len(name) == 0 {
		return nil, xerrors.New("name must not be empty")
	}
	if len(name) > nameMaxLength {
		return nil, xerrors.Errorf("name must less than %d: %s", nameMaxLength, name)
	}
	if len(profile) == 0 {
		return nil, xerrors.New("profile must not be empty")
	}
	if len(profile) > profileMaxLength {
		return nil, xerrors.Errorf("profile must less than %d: %s", profileMaxLength, profile)
	}

	userId, err := NewUserId()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	emailIns, err := NewEmail(email)
	if err != nil {
		return nil, err
	}
	user := &User{
		UserId:    userId,
		Name:      name,
		Email:     emailIns,
		Password:  password,
		Profile:   profile,
		CreatedAt: now,
	}

	return user, nil
}
