package repoimpl

import (
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

func TestUserRepo_Create(t *testing.T) {
	const (
		name     = "テスト"
		email    = "test@gmail.com"
		password = "abcd12341231"
		profile  = "プログラマーです"
		from     = "2020-03-03"
		to       = "2020-04-04"
		detail   = "DDD"
	)
	userID, err := userdm.NewUserID()
	if err != nil {
		t.Errorf("failed to NewUserID: %v", err)
	}
	emailIns, err := userdm.NewEmail(email)
	if err != nil {
		t.Errorf("failed to NewEmail: %v", err)
	}
	user, err := userdm.NewUser(userID, name, emailIns, password, profile)
	if err != nil {
		t.Errorf("failed to NewUser: %v", err)
	}

	userRepository := NewUserRepositoryImpl(NewDB())
	createdUser, err := userRepository.Create(user)
	if err != nil {
		t.Errorf("failed to userRepository.Create: %v", err)
	}
	selectedUser, err := userRepository.FindByID(createdUser.UserID)
	if err != nil {
		t.Errorf("failed to FindByID: %v", err)
	}
	if !userdm.UserID.Equals(createdUser.UserID, selectedUser.UserID) {
		t.Errorf("failed to CreateUser")
	}
}
