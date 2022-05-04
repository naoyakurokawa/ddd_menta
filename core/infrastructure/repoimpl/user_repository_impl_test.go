package repoimpl

import (
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
)

func TestUserRepo_Create(t *testing.T) {
	const (
		name     = "テスト"
		email    = "test@gmail.com"
		password = "abcd12341231"
		profile  = "プログラマーです"
	)

	userID := userdm.NewUserID()
	userCareers := make([]userdm.UserCareer, 2)
	userCareerID1 := userdm.NewUserCareerID()
	userCareerID2 := userdm.NewUserCareerID()
	userCareer1, err := userdm.NewUserCareer(userCareerID1, "2013-06-02 15:04:05", "2013-06-02 15:04:05", "PHPエンジニア")
	if err != nil {
		t.Errorf("failed to NewUserCareer: %v", err)
		return
	}
	userCareer2, err := userdm.NewUserCareer(userCareerID2, "2013-06-02 15:04:05", "2013-06-02 15:04:05", "Goエンジニア")
	if err != nil {
		t.Errorf("failed to NewUserCareer: %v", err)
		return
	}
	userCareers[0] = *userCareer1
	userCareers[1] = *userCareer2

	userSkills := make([]userdm.UserSkill, 1)
	userSkillID := userdm.NewUserSkillID()
	if err != nil {
		t.Errorf("failed to NewUserSkillID: %v", err)
		return
	}
	experienceYears, err := userdm.NewExperienceYears(1)
	if err != nil {
		t.Errorf("failed to NewExperienceYears: %v", err)
		return
	}
	userSkill, err := userdm.NewUserSkill(userSkillID, "Golang", 5, experienceYears)
	if err != nil {
		t.Errorf("failed to NewUserSkill: %v", err)
		return
	}
	userSkills[0] = *userSkill

	emailIns, err := userdm.NewEmail(email)
	if err != nil {
		t.Errorf("failed to NewEmail: %v", err)
		return
	}

	user, err := userdm.NewUser(
		userID,
		name,
		emailIns,
		password,
		profile,
		userCareers,
		userSkills,
	)
	if err != nil {
		t.Errorf("failed to NewUser: %v", err)
		return
	}

	userRepository := NewUserRepositoryImpl(db.NewDB())
	err = userRepository.Create(user)
	if err != nil {
		t.Errorf("failed to userRepository.Create: %v", err)
	}
	selectedUser, err := userRepository.FetchById(user.UserID())
	if err != nil {
		t.Errorf("failed to FindByID: %v", err)
	}
	if !user.UserID().Equals(selectedUser.UserID()) {
		t.Errorf("failed to CreateUser")
	}
}
