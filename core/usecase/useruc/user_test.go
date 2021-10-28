package useruc

import (
	"testing"

	// "github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	// "github.com/naoyakurokawa/ddd_menta/core/infrastructure/repoimpl"
	"github.com/golang/mock/gomock"

	mock "github.com/naoyakurokawa/ddd_menta/core/usecase/useruc/mock_useruc"
	// "time"
)

func TestCreate(t *testing.T) {
	const (
		name     = "テスト"
		email    = "test@gmail.com"
		password = "abcd12341231"
		profile  = "プログラマーです"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserUsecase := mock.NewMockUserUsecase(ctrl)

	mockUserUsecase.EXPECT().Create(name, email, password, profile)
	_, err := mockUserUsecase.Create(name, email, password, profile)

	if err != nil {
		t.Errorf("failed to userRepository.Create: %v", err)
	}
}
