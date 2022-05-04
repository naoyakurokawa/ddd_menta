package userdm

import (
	"strings"
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
)

func TestNewUser(t *testing.T) {
	const (
		name     = "テストユーザー"
		email    = "test@gmail.com"
		password = "test12345678"
		profile  = "エンジニアです"
	)
	userID := NewUserID()
	emailIns, err := sharedvo.NewEmail(email)
	if err != nil {
		t.Errorf("failed to NewEmail: %v", err)
		return
	}
	passwordIns, err := sharedvo.NewPassword(password)
	if err != nil {
		t.Errorf("failed to NewPassword: %v", err)
		return
	}
	userCareers := []UserCareer{}
	userSkills := []UserSkill{}

	t.Run("UserIDが空", func(t *testing.T) {
		BlankUserID := UserID("")
		_, err := NewUser(BlankUserID, name, emailIns, passwordIns, profile, userCareers, userSkills)
		if err == nil {
			t.Errorf("failed to UserID empty validation")
		}
	})

	t.Run("Nameが空", func(t *testing.T) {
		blankName := ""
		_, err := NewUser(userID, blankName, emailIns, passwordIns, profile, userCareers, userSkills)
		if err == nil {
			t.Errorf("failed to Name empty validation")
		}
	})

	t.Run("Nameが最大文字数超過", func(t *testing.T) {
		nameOver := strings.Repeat("a", 256)
		_, err := NewUser(userID, nameOver, emailIns, passwordIns, profile, userCareers, userSkills)
		if err == nil {
			t.Errorf("failed to Name maxlength validation: %v", nameOver)
		}
	})

	t.Run("Emailが空", func(t *testing.T) {
		blankEmail := sharedvo.Email("")
		_, err := NewUser(userID, name, blankEmail, passwordIns, profile, userCareers, userSkills)
		if err == nil {
			t.Errorf("failed to Email empty validation")
		}
	})

	t.Run("Passwordが空", func(t *testing.T) {
		blankPassword := sharedvo.Password("")
		_, err := NewUser(userID, name, email, blankPassword, profile, userCareers, userSkills)
		if err == nil {
			t.Errorf("failed to Password empty validation")
		}
	})

	t.Run("Profileが空", func(t *testing.T) {
		blankProfile := ""
		_, err := NewUser(userID, name, email, passwordIns, blankProfile, userCareers, userSkills)
		if err == nil {
			t.Errorf("failed to Profile empty validation")
		}
	})

	t.Run("Profilが最大文字数超過", func(t *testing.T) {
		profileOver := strings.Repeat("a", 2001)
		_, err := NewUser(userID, name, email, passwordIns, profileOver, userCareers, userSkills)
		if err == nil {
			t.Errorf("failed to profile maxlength validation: %v", profileOver)
		}
	})
}
