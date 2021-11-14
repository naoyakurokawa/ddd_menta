package userdm

import (
	"time"

	"golang.org/x/xerrors"
)

type User struct {
	UserID     UserID
	Name       string
	Email      Email
	Password   string
	Profile    string
	CreatedAt  time.Time
	UserCareers []UserCareer
}

const nameMaxLength = 255
const profileMaxLength = 2000

// NewUser userのコンストラクタ
func NewUser(userID UserID, name string, email Email, password string, profile string, userCareers []UserCareer) (*User, error) {
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

	now := time.Now()

	user := &User{
		UserID:     userID,
		Name:       name,
		Email:      email,
		Password:   password,
		Profile:    profile,
		CreatedAt:  now,
		UserCareers: userCareers,
	}

	return user, nil
}
