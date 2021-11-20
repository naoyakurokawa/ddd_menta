package userdm

import (
	"strings"
	"testing"
)

func TestNewUser(t *testing.T) {
	const (
		name     = "テストユーザー"
		email    = "test@gmail.com"
		password = "test1234"
		profile  = "エンジニアです"
	)
	userID, err := NewUserID()
	if err != nil {
		t.Errorf("failed to NewUserID: %v", err)
	}
	emailIns, err := NewEmail(email)
	if err != nil {
		t.Errorf("failed to NewUserID: %v", err)
	}
	userCareers := []UserCareer{}

	t.Run("UserIDが空", func(t *testing.T) {
		BlankUserID := UserID("")
		_, err := NewUser(BlankUserID, name, emailIns, password, profile, userCareers)
		if err == nil {
			t.Errorf("failed to UserID empty validation: %v", err)
		}
	})

	t.Run("Nameが空", func(t *testing.T) {
		blankName := ""
		_, err := NewUser(userID, blankName, emailIns, password, profile, userCareers)
		if err == nil {
			t.Errorf("failed to Name empty validation: %v", err)
		}
	})

	t.Run("Nameが最大文字数超過", func(t *testing.T) {
		nameOver := strings.Repeat("a", 256)
		_, err := NewUser(userID, nameOver, emailIns, password, profile, userCareers)
		if err == nil {
			t.Errorf("failed to Name maxlength validation: %v", err)
		}
	})

	t.Run("Emailが空", func(t *testing.T) {
		blankEmail := Email("")
		_, err := NewUser(userID, name, blankEmail, password, profile, userCareers)
		if err == nil {
			t.Errorf("failed to Email empty validation: %v", err)
		}
	})

	t.Run("Passwordが空", func(t *testing.T) {
		blankPassword := ""
		_, err := NewUser(userID, name, email, blankPassword, profile, userCareers)
		if err == nil {
			t.Errorf("failed to Password empty validation: %v", err)
		}
	})

	t.Run("Profileが空", func(t *testing.T) {
		blankProfile := ""
		_, err := NewUser(userID, name, email, password, blankProfile, userCareers)
		if err == nil {
			t.Errorf("failed to Profile empty validation: %v", err)
		}
	})

	t.Run("Profilが最大文字数超過", func(t *testing.T) {
		profileOver := strings.Repeat("a", 2001)
		_, err := NewUser(userID, name, email, password, profileOver, userCareers)
		if err == nil {
			t.Errorf("failed to profile maxlength validation: %v", err)
		}
	})
}
