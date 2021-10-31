package useruc

import (
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	// "github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/golang/mock/gomock"

	mock "github.com/naoyakurokawa/ddd_menta/core/domain/userdm/mock_userdm"
	// "time"
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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserUsecase := mock.NewMockUserRepository(ctrl)

	mockUserUsecase.EXPECT().Create(user)
	_, err = mockUserUsecase.Create(user)

	if err != nil {
		t.Errorf("failed to userRepository.Create: %v", err)
	}
}
