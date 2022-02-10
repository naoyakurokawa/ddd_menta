package mentordm

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMentor(t *testing.T) {

	t.Run("titleが空の場合_エラーとなること", func(t *testing.T) {
		setup()
		blankTitle := ""
		_, err := NewMentor(mp.userID, mp.mentorID, blankTitle, mp.mainImg, mp.subImg, mp.category, mp.detial)
		require.Error(t, err)
	})

	t.Run("titleが255文字を超えるの場合_エラーとなること", func(t *testing.T) {
		setup()
		overTitle := strings.Repeat("a", 256)
		_, err := NewMentor(mp.userID, mp.mentorID, overTitle, mp.mainImg, mp.subImg, mp.category, mp.detial)
		require.Error(t, err)
	})

	t.Run("detailが空の場合_エラーとなること", func(t *testing.T) {
		setup()
		blankDetail := ""
		_, err := NewMentor(mp.userID, mp.mentorID, mp.title, mp.mainImg, mp.subImg, mp.category, blankDetail)
		require.Error(t, err)
	})

	t.Run("detailが2000文字を超える場合_エラーとなること", func(t *testing.T) {
		setup()
		overDetail := strings.Repeat("a", 2001)
		_, err := NewMentor(mp.userID, mp.mentorID, mp.title, mp.mainImg, mp.subImg, mp.category, overDetail)
		require.Error(t, err)
	})
}

func TestAddMentorSkill(t *testing.T) {
	t.Run("生成したMentorSkillがMentorに追加されていること", func(t *testing.T) {
		setup()

		// 期待値
		expectedTag := "Golang"
		expectedAssessment := uint16(5)
		expectedExperienceYears, err := NewExperienceYears(5)
		require.NoError(t, err)

		// メンター作成
		m, err := NewMentor(mp.userID, mp.mentorID, mp.title, mp.mainImg, mp.subImg, mp.category, mp.detial)
		require.NoError(t, err)

		// メンタースキル追加実行
		actual, err := m.AddMentorSkill(mp.mentorTag, mp.mentorAssessment, mp.mentorExperienceYears)
		require.NoError(t, err)

		// 検証
		assert.Equal(t, expectedTag, actual.MentorSkills()[0].tag)
		assert.Equal(t, expectedAssessment, actual.MentorSkills()[0].assessment)
		assert.Equal(t, expectedExperienceYears, actual.MentorSkills()[0].experienceYears)
	})
}

func TestAddPlan(t *testing.T) {
	t.Run("生成したPlanがMentorに追加されていること", func(t *testing.T) {
		setup()

		// 期待値
		expectedTitle := "DDDのメンタリング"
		expectedCategory := "設計"
		expectedTag := "DDD"
		expectedDetial := "DDDの設計手法を学べます"
		expectedPlanType, err := NewPlanType(2)
		require.NoError(t, err)
		expectedPrice := uint16(1000)
		expectedPlanStatus, err := NewPlanStatus(1)
		require.NoError(t, err)

		// メンター作成
		m, err := NewMentor(mp.userID, mp.mentorID, mp.title, mp.mainImg, mp.subImg, mp.category, mp.detial)
		require.NoError(t, err)

		// プラン追加実行
		actual, err := m.AddPlan(mp.planTitle, mp.planCategory, mp.planTag, mp.planDetial, mp.planType, mp.planPrice, mp.planStatus)
		require.NoError(t, err)

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
