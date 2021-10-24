package useruc

import (
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
)

const (
	Name     = "テスト"
	Email    = "test@gmail.com"
	Password = "abcd12341231"
	Profile  = "プログラマーです"
)

// Create userを保存するときのユースケース
func TestCreate(t *testing.T) {
	userID, err := userdm.NewUserId()
	if err != nil {
		t.Errorf("failed to NewUserId: %v", err)
	}
	emailIns, err := userdm.NewEmail(Email)
	if err != nil {
		t.Errorf("failed to NewEmail: %v", err)
	}
	user, err := userdm.NewUser(userID, Name, emailIns, Password, Profile)
	if err != nil {
		t.Errorf("failed to NewUser: %v", err)
	}
	userRepository := repoimpl.NewUserRepositoryImpl(repoimpl.NewDB())

	_, err = userRepository.Create(user)
	if err != nil {
		t.Errorf("failed to userRepository.Create: %v", err)
	}
}
