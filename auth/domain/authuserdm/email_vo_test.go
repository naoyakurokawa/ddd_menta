package authuserdm

import (
	"strings"
	"testing"
)

func TestNewEmail(t *testing.T) {
	t.Run("Emailが空", func(t *testing.T) {
		emailBlank := ""
		_, err := NewEmail(emailBlank)
		if err == nil {
			t.Errorf("failed to email empty validation")
		}
	})

	t.Run("Emailが最大文字数超過", func(t *testing.T) {
		emailOver := strings.Repeat("a", 246)
		emailOver += "@gmail.com"
		_, err := NewEmail(emailOver)
		if err == nil {
			t.Errorf("failed to email max length validation: %v", emailOver)
		}
	})

	t.Run("Emailがフォーマットエラー", func(t *testing.T) {
		emailNotSuitableFormat := "aaaa"
		_, err := NewEmail(emailNotSuitableFormat)

		if err == nil {
			t.Errorf("failed to email max length validation: %v", emailNotSuitableFormat)
		}
	})
}
