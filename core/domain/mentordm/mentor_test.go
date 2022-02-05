package mentordm

import (
	"strings"
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

func TestNewMentor(t *testing.T) {
	const (
		title    = "プログラミング全般のメンタリング"
		mainImg  = "/main.jpg"
		subImg   = "/sub.jpg"
		category = "プログライミング"
		detial   = "設計・開発・テストの一覧をサポートできます"
	)
	userID, err := userdm.NewUserID()
	if err != nil {
		t.Errorf("failed to NewUserID: %v", err)
		return
	}
	mentorID := NewMentorID()

	// mentorSkills := []MentorSkill{}
	// plans := []Plan{}
	t.Run("titleが空の場合_エラーとなること", func(t *testing.T) {
		blankTitle := ""
		_, err := NewMentor(userID, mentorID, blankTitle, mainImg, subImg, category, detial)
		if err == nil {
			t.Errorf("failed to title blank validation")
		}
	})

	t.Run("titleが255文字を超えるの場合_エラーとなること", func(t *testing.T) {
		overTitle := strings.Repeat("a", 256)
		_, err := NewMentor(userID, mentorID, overTitle, mainImg, subImg, category, detial)
		if err == nil {
			t.Errorf("failed to title maxlength validation: %v", overTitle)
		}
	})

	t.Run("detailが空の場合_エラーとなること", func(t *testing.T) {
		blankDetail := ""
		_, err := NewMentor(userID, mentorID, title, mainImg, subImg, category, blankDetail)
		if err == nil {
			t.Errorf("failed to detail blank validation")
		}
	})

	t.Run("detailが2000文字を超える場合_エラーとなること", func(t *testing.T) {
		overDetail := strings.Repeat("a", 2001)
		_, err := NewMentor(userID, mentorID, title, mainImg, subImg, category, overDetail)
		if err == nil {
			t.Errorf("failed to title maxlength validation: %v", overDetail)
		}
	})
}
