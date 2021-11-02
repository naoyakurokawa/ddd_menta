package repoimpl

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"testing"
)

func TestCreate(t *testing.T) {
	const (
		name     = "テスト"
		email    = "test@gmail.com"
		password = "abcd12341231"
		profile  = "プログラマーです"
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
	user2, err := userRepository.Create(user)
	user3 := userRepository.FindByID(user2);
	if user2.UserID != user3.UserID {
		t.Errorf("failed to CreateUser")
	}
}
