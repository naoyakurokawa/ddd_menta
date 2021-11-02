package userdm

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByID(user *User) (*User)
}
