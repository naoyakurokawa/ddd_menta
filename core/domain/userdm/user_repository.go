package userdm

import "github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"

type UserRepository interface {
	Create(user *User) error
	FetchById(userID UserID) (*User, error)
	FetchByEmail(email sharedvo.Email) (*User, error)
}
