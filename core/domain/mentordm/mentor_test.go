package mentordm

import (
	"strings"
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	mentorSkills []MentorSkill
	plans        []Plan
)

func TestNewMentor(t *testing.T) {
	for _, td := range []struct {
		title        string
		userID       userdm.UserID
		mentorTitle  string
		mainImg      string
		subImg       string
		category     string
		detail       string
		mentorSkills []MentorSkill
		plans        []Plan
	}{
		{
			title:        "titleが空の場合_エラーとなること",
			userID:       mp.userID,
			mentorTitle:  "",
			mainImg:      mp.mainImg,
			subImg:       mp.subImg,
			category:     mp.category,
			detail:       mp.detial,
			mentorSkills: mentorSkills,
			plans:        plans,
		},
		{
			title:        "titleが255文字を超えるの場合_エラーとなること",
			userID:       mp.userID,
			mentorTitle:  strings.Repeat("a", 256),
			mainImg:      mp.mainImg,
			subImg:       mp.subImg,
			category:     mp.category,
			detail:       mp.detial,
			mentorSkills: mentorSkills,
			plans:        plans,
		},
		{
			title:        "detailが空の場合_エラーとなること",
			userID:       mp.userID,
			mentorTitle:  mp.title,
			mainImg:      mp.mainImg,
			subImg:       mp.subImg,
			category:     mp.category,
			detail:       "",
			mentorSkills: mentorSkills,
			plans:        plans,
		},
		{
			title:        "detailが2000文字を超える場合_エラーとなること",
			userID:       mp.userID,
			mentorTitle:  mp.title,
			mainImg:      mp.mainImg,
			subImg:       mp.subImg,
			category:     mp.category,
			detail:       strings.Repeat("a", 2001),
			mentorSkills: mentorSkills,
			plans:        plans,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			_, err := NewMentor(mp.mentorID,
				td.userID,
				td.mentorTitle,
				td.mainImg,
				td.subImg,
				td.category,
				td.detail,
				td.mentorSkills,
				td.plans,
			)
			require.Error(t, err)
		})
	}
}

func TestAddMentorSkill(t *testing.T) {
	t.Run("生成したMentorSkillがMentorに追加されていること", func(t *testing.T) {
		err := setup()
		require.NoError(t, err)

		// 期待値
		expectedTag := "Golang"
		expectedAssessment := uint16(5)
		expectedExperienceYears, err := NewExperienceYears(5)
		require.NoError(t, err)

		// メンター作成
		m, err := NewMentor(
			mp.mentorID,
			mp.userID,
			mp.title,
			mp.mainImg,
			mp.subImg,
			mp.category,
			mp.detial,
			mentorSkills,
			plans,
		)
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
		err := setup()
		require.NoError(t, err)

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
		m, err := NewMentor(
			mp.mentorID,
			mp.userID,
			mp.title,
			mp.mainImg,
			mp.subImg,
			mp.category,
			mp.detial,
			mentorSkills,
			plans,
		)
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
