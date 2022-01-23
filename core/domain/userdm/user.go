package userdm

import (
	"time"
	"unicode/utf8"

	"golang.org/x/xerrors"
)

type User struct {
	userID      UserID
	name        string
	email       Email
	password    Password
	profile     string
	createdAt   time.Time
	userCareers []UserCareer
	userSkills  []UserSkill
}

const nameMaxLength = 255
const profileMaxLength = 2000

// NewUser userのコンストラクタ
func NewUser(userID UserID, name string, email Email, password Password, profile string, userCareers []UserCareer, userSkills []UserSkill) (*User, error) {
	//入力データチェック
	if len(userID) == 0 {
		return nil, xerrors.New("userID must not be empty")
	}
	if len(name) == 0 {
		return nil, xerrors.New("name must not be empty")
	}
	if utf8.RuneCountInString(name) > nameMaxLength {
		return nil, xerrors.Errorf("name must less than %d: %s", nameMaxLength, name)
	}
	if len(email) == 0 {
		return nil, xerrors.New("email must not be empty")
	}
	if len(password) == 0 {
		return nil, xerrors.New("password must not be empty")
	}
	if len(profile) == 0 {
		return nil, xerrors.New("profile must not be empty")
	}
	if utf8.RuneCountInString(profile) > profileMaxLength {
		return nil, xerrors.Errorf("profile must less than %d: %s", profileMaxLength, profile)
	}

	now := time.Now()

	user := &User{
		userID:      userID,
		name:        name,
		email:       email,
		password:    password,
		profile:     profile,
		createdAt:   now,
		userCareers: userCareers,
		userSkills:  userSkills,
	}

	return user, nil
}

func (u *User) UserID() UserID {
	return u.userID
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) Password() Password {
	return u.password
}

func (u *User) Profile() string {
	return u.profile
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UserCareers() []UserCareer {
	return u.userCareers
}

func (u *User) UserSkills() []UserSkill {
	return u.userSkills
}
