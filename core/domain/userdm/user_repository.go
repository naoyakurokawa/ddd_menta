package userdm

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByID(userID UserID) (*User, error)
	CreateUserCareers(userCareers *UserCareers) (*UserCareers, error)
}
