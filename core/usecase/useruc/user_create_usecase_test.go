package useruc_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/usecase/useruc"

	mock "github.com/naoyakurokawa/ddd_menta/core/domain/userdm/mock_userdm"
)

func TestCreate(t *testing.T) {
	const (
		name     = "テスト"
		email    = "test@gmail.com"
		password = "abcd12341231"
		profile  = "プログラマーです"
	)
	from := []string{"2013-06-02 15:04:05", "2013-06-02 15:04:05"}
	to := []string{"2013-06-02 15:04:05", "2013-06-02 15:04:05"}
	detail := []string{"PHPエンジニア", "Goエンジニア"}

	userID, err := userdm.NewUserID()
	if err != nil {
		t.Errorf("failed to NewUserID: %v", err)
	}
	// userCareers := []userdm.UserCareer{}
	userCareers := make([]userdm.UserCareer, 2)
	userCareerID1, err := userdm.NewUserCareerID()
	if err != nil {
		t.Errorf("failed to NewUserCareerID: %v", err)
	}
	userCareerID2, err := userdm.NewUserCareerID()
	if err != nil {
		t.Errorf("failed to NewUserCareerID: %v", err)
	}
	userCareer1, err := userdm.NewUserCareer(userCareerID1, userID, "2013-06-02 15:04:05", "2013-06-02 15:04:05", "PHPエンジニア")
	if err != nil {
		t.Errorf("failed to NewUserCareer: %v", err)
	}
	userCareer2, err := userdm.NewUserCareer(userCareerID2, userID, "2013-06-02 15:04:05", "2013-06-02 15:04:05", "Goエンジニア")
	if err != nil {
		t.Errorf("failed to NewUserCareer: %v", err)
	}
	userCareers[0] = *userCareer1
	userCareers[1] = *userCareer2

	emailIns, err := userdm.NewEmail(email)
	if err != nil {
		t.Errorf("failed to NewEmail: %v", err)
	}
	user, err := userdm.NewUser(userID, name, emailIns, password, profile, userCareers)
	if err != nil {
		t.Errorf("failed to NewUser: %v", err)
	}

	ctrl := gomock.NewController(t)

	mockUserRepository := mock.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().Create(gomock.Any()).Return(user, nil)
	userusecase := useruc.NewUserCreateUsecase(mockUserRepository)
	_, err = userusecase.Create(name, email, password, profile, from, to, detail)

	if err != nil {
		t.Errorf("failed to userRepository.Create: %v", err)
	}
}
