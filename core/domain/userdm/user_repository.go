package userdm

type UserRepository interface {
	Create(user *User) error
	FetchById(userID UserID) (*User, error)
	FetchByEmail(email Email) (*User, error)
}
