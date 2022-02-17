package userdm

import (
	"strings"
	"testing"
)

func TestNewUserCareer(t *testing.T) {
	const (
		from   = "2006-01-02 15:04:05"
		to     = "2006-01-02 15:04:05"
		detail = "プログラマー"
	)
	userID := NewUserID()
	userCareerID, err := NewUserCareerID()
	if err != nil {
		t.Errorf("failed to NewUserCareerID: %v", err)
		return
	}

	t.Run("UserCareerIDが空", func(t *testing.T) {
		blankUserCareerID := UserCareerID("")
		_, err := NewUserCareer(blankUserCareerID, userID, from, to, detail)
		if err == nil {
			t.Errorf("failed to NewUserCareerID empty validation: %v", blankUserCareerID)
		}
	})

	t.Run("UserIDが空", func(t *testing.T) {
		blankUserID := UserID("")
		_, err := NewUserCareer(userCareerID, blankUserID, from, to, detail)
		if err == nil {
			t.Errorf("failed to UserID empty validation: %v", blankUserID)
		}
	})

	t.Run("Fromが空", func(t *testing.T) {
		blankFrom := ""
		_, err := NewUserCareer(userCareerID, userID, blankFrom, to, detail)
		if err == nil {
			t.Errorf("failed to From empty validation: %v", blankFrom)
		}
	})

	t.Run("Fromが1969年以前", func(t *testing.T) {
		fromNotSuitableFormat := "1969-01-02 15:04:05"
		_, err := NewUserCareer(userCareerID, userID, fromNotSuitableFormat, to, detail)
		if err == nil {
			t.Errorf("failed to From datecheck validation: %v", fromNotSuitableFormat)
		}
	})

	t.Run("Toが空", func(t *testing.T) {
		blankTo := ""
		_, err := NewUserCareer(userCareerID, userID, from, blankTo, detail)
		if err == nil {
			t.Errorf("failed to To empty validation: %v", blankTo)
		}
	})

	t.Run("Toが1969年以前", func(t *testing.T) {
		toNotSuitableFormat := "1969-01-02 15:04:05"
		_, err := NewUserCareer(userCareerID, userID, from, toNotSuitableFormat, detail)
		if err == nil {
			t.Errorf("failed to From datecheck validation: %v", toNotSuitableFormat)
		}
	})

	t.Run("FromがTo以前", func(t *testing.T) {
		failTo := "1970-01-02 15:04:05"
		failFrom := "1971-01-02 15:04:05"
		_, err := NewUserCareer(userCareerID, userID, failFrom, failTo, detail)
		if err == nil {
			t.Errorf("failed to From and To datecheck validation: to: %v from: %v", failTo, failFrom)
		}
	})

	t.Run("Detailが空", func(t *testing.T) {
		blankDetail := ""
		_, err := NewUserCareer(userCareerID, userID, from, to, blankDetail)
		if err == nil {
			t.Errorf("failed to Detail empty validation: %v", blankDetail)
		}
	})

	t.Run("Detailが最大文字数超過", func(t *testing.T) {
		detailOver := strings.Repeat("a", 1001)
		_, err := NewUserCareer(userCareerID, userID, from, to, detailOver)
		if err == nil {
			t.Errorf("failed to Detail maxlength validation: %v", detailOver)
		}
	})
}
