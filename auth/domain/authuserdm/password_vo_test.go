package authuserdm

import (
	"strings"
	"testing"
)

func TestNewPassword(t *testing.T) {
	t.Run("Passwordが空", func(t *testing.T) {
		passwordBlank := ""
		_, err := NewPassword(passwordBlank)
		if err == nil {
			t.Errorf("failed to password empty validation")
		}
	})

	t.Run("Passwordが最小文字未満", func(t *testing.T) {
		belowPassword := strings.Repeat("a", 5)
		belowPassword += strings.Repeat("1", 5)
		_, err := NewPassword(belowPassword)
		if err == nil {
			t.Errorf("failed to password min length validation: %v", belowPassword)
		}
	})

	t.Run("passwordに数字が含まれていない", func(t *testing.T) {
		PasswordNotSuitableFormat := strings.Repeat("a", 12)
		_, err := NewPassword(PasswordNotSuitableFormat)

		if err == nil {
			t.Errorf("failed to password include number validation: %v", PasswordNotSuitableFormat)
		}
	})

	t.Run("passwordに英字が含まれていない", func(t *testing.T) {
		PasswordNotSuitableFormat := strings.Repeat("1", 12)
		_, err := NewPassword(PasswordNotSuitableFormat)

		if err == nil {
			t.Errorf("failed to password include alphabet validation: %v", PasswordNotSuitableFormat)
		}
	})

	t.Run("passwordにスペースが含まれている", func(t *testing.T) {
		PasswordNotSuitableFormat := "aaa 1111 aaaa 1111"
		_, err := NewPassword(PasswordNotSuitableFormat)

		if err == nil {
			t.Errorf("failed to password notinclude space validation: %v", PasswordNotSuitableFormat)
		}
	})
}
