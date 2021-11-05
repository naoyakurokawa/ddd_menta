package userdm

import (
	"testing"
	"strings"
)

func TestNewEmail(t *testing.T) {
	t.Run("Emailが空", func(t *testing.T) {
		emailBlank := ""
		_, err := NewEmail(emailBlank)
		if err == nil {
			t.Errorf("failed to email empty validation: %v", err)
		}
	})

	t.Run("Emailが最大文字数超過", func(t *testing.T) {
		emailOver := strings.Repeat("a", 246)
		emailOver += "@gmail.com"
		_, err := NewEmail(emailOver)
		if err == nil {
			t.Errorf("failed to email max length validation: %v", err)
		}
	})

	t.Run("Emailがフォーマットエラー", func(t *testing.T) {
		emailNotSuitableFormat := "aaaa"
		_, err := NewEmail(emailNotSuitableFormat)
	
		if err == nil {
			t.Errorf("failed to email max length validation: %v", err)
		}
	})
}
