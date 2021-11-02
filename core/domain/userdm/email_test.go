package userdm

import (
	"testing"
)

func TestNewEmail(t *testing.T) {
	//given:前提条件
	emailBlank := ""
	emailOver := ""
	emailNotSuitableFormat := "aaaa"
	for i := 0; i < 300; i++ {
		emailOver += "a"
	}
	//when：操作
	_, err := NewEmail(emailBlank)
	//then：結果
	if err == nil {
		t.Errorf("failed to email empty validation: %v", err)
	}
	//when：操作
	_, err = NewEmail(emailOver)
	//then：結果
	if err == nil {
		t.Errorf("failed to email max length validation: %v", err)
	}
	//when：操作
	_, err = NewEmail(emailNotSuitableFormat)
	//then：結果
	if err == nil {
		t.Errorf("failed to email max length validation: %v", err)
	}
}
