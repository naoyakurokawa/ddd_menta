package userdm

import (
	"log"
	"time"
)

type User struct {
	UserId    string
	Name      string
	Email     string
	Password  string
	Profile   string
	CreatedAt time.Time
}

// NewUser userのコンストラクタ
func NewUser(name string, email string, password string, profile string) (*User, error) {
	user_id, err := NewUserId()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	now := time.Now()
	user := &User{
		UserId:    user_id.UserId,
		Name:      name,
		Email:     email,
		Password:  password,
		Profile:   profile,
		CreatedAt: now,
	}

	return user, nil
}
