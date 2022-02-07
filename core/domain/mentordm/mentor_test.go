package mentordm

import (
	"strings"
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/stretchr/testify/assert"
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

func TestAddMentorSkill(t *testing.T) {
	t.Run("生成したMentorSkillがMentorに追加されていること", func(t *testing.T) {
		// 期待値
		expectedTag := "Golang"
		expectedAssessment := uint16(5)
		expectedExperienceYears, _ := NewExperienceYears(5)

		// メンター作成テストデータ
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

		// メンター作成
		m, _ := NewMentor(userID, mentorID, title, mainImg, subImg, category, detial)

		// メンタースキル追加実行
		tag := "Golang"
		assessment := uint16(5)
		experienceYears := uint16(5)
		actual, err := m.AddMentorSkill(tag, assessment, experienceYears)
		if err != nil {
			t.Errorf("failed to AddMentorSkill")
		}

		// 検証
		assert.Equal(t, expectedTag, actual.MentorSkills()[0].tag)
		assert.Equal(t, expectedAssessment, actual.MentorSkills()[0].assessment)
		assert.Equal(t, expectedExperienceYears, actual.MentorSkills()[0].experienceYears)
	})
}

func TestAddPlan(t *testing.T) {
	t.Run("生成したPlanがMentorに追加されていること", func(t *testing.T) {
		// 期待値
		expectedTitle := "DDDのメンタリング"
		expectedCategory := "設計"
		expectedTag := "DDD"
		expectedDetial := "DDDの設計手法を学べます"
		expectedPlanType, _ := NewPlanType(2)
		expectedPrice := uint16(1000)
		expectedPlanStatus, _ := NewPlanStatus(1)

		// メンター作成テストデータ
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

		// メンター作成
		m, _ := NewMentor(userID, mentorID, title, mainImg, subImg, category, detial)

		// プラン追加実行
		planTitle := "DDDのメンタリング"
		planCategory := "設計"
		tag := "DDD"
		detail := "DDDの設計手法を学べます"
		planType, _ := NewPlanType(uint16(2))
		price := uint16(1000)
		planStatus, _ := NewPlanStatus(uint16(1))
		actual, err := m.AddPlan(planTitle, planCategory, tag, detail, planType, price, planStatus)
		if err != nil {
			t.Errorf("failed to AddPlan")
		}

		// 検証
		assert.Equal(t, expectedTitle, actual.Plans()[0].title)
		assert.Equal(t, expectedCategory, actual.Plans()[0].category)
		assert.Equal(t, expectedTag, actual.Plans()[0].tag)
		assert.Equal(t, expectedDetial, actual.Plans()[0].detial)
		assert.Equal(t, expectedPlanType, actual.Plans()[0].planType)
		assert.Equal(t, expectedPrice, actual.Plans()[0].price)
		assert.Equal(t, expectedPlanStatus, actual.Plans()[0].planStatus)
	})
}
