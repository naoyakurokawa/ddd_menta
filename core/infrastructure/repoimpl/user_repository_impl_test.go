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
	)
	// from := []string{"2013-06-02 15:04:05", "2013-06-02 15:04:05"}
	// to := []string{"2013-06-02 15:04:05", "2013-06-02 15:04:05"}
	// detail := []string{"PHPエンジニア", "Goエンジニア"}
	userID, err := userdm.NewUserID()
	if err != nil {
		t.Errorf("failed to NewUserID: %v", err)
	}
	userCareers := []userdm.UserCareer{}
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
	userCareers = append(userCareers, *userCareer1, *userCareer2)
	emailIns, err := userdm.NewEmail(email)
	if err != nil {
		t.Errorf("failed to NewEmail: %v", err)
	}
	user, err := userdm.NewUser(userID, name, emailIns, password, profile, userCareers)
	if err != nil {
		t.Errorf("failed to NewUser: %v", err)
	}
	userRepository := NewUserRepositoryImpl(NewDB())
	createdUser, err := userRepository.Create(user)
	if err != nil {
		t.Errorf("failed to userRepository.Create: %v", err)
	}
	selectedUser, err := userRepository.FindByID(createdUser.UserID())
	if err != nil {
		t.Errorf("failed to FindByID: %v", err)
	}
	if !userdm.UserID.Equals(createdUser.UserID(), selectedUser.UserID()) {
		t.Errorf("failed to CreateUser")
	}
}
