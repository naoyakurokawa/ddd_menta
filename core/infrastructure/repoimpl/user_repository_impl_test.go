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

	userID, err := userdm.NewUserID()
	if err != nil {
		t.Errorf("failed to NewUserID: %v", err)
	}

	userCareers := []userdm.UserCareer{}
	userCareerID1, err := userdm.NewUserCareerID()
	if err != nil {
		t.Errorf("failed to NewUserCareerID: %v", err)
		return
	}
	userCareerID2, err := userdm.NewUserCareerID()
	if err != nil {
		t.Errorf("failed to NewUserCareerID: %v", err)
		return
	}
	userCareer1, err := userdm.NewUserCareer(userCareerID1, userID, "2013-06-02 15:04:05", "2013-06-02 15:04:05", "PHPエンジニア")
	if err != nil {
		t.Errorf("failed to NewUserCareer: %v", err)
		return
	}
	userCareer2, err := userdm.NewUserCareer(userCareerID2, userID, "2013-06-02 15:04:05", "2013-06-02 15:04:05", "Goエンジニア")
	if err != nil {
		t.Errorf("failed to NewUserCareer: %v", err)
		return
	}
	userCareers = append(userCareers, *userCareer1, *userCareer2)

	emailIns, err := userdm.NewEmail(email)
	if err != nil {
		t.Errorf("failed to NewEmail: %v", err)
		return
	}

	user, err := userdm.NewUser(userID, name, emailIns, password, profile, userCareers)
	if err != nil {
		t.Errorf("failed to NewUser: %v", err)
		return
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
