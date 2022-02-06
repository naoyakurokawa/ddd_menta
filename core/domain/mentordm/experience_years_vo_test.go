package mentordm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExperienceYears(t *testing.T) {
	t.Run("ExperienceYearsが1のとき_エラーが発生しないこと", func(t *testing.T) {
		_, err := NewExperienceYears(1)
		if err != nil {
			t.Errorf("failed to NewExperienceYears")
		}
	})

	t.Run("ExperienceYearsが6のとき_エラーが発生すること", func(t *testing.T) {
		_, err := NewExperienceYears(6)
		if err == nil {
			t.Errorf("failed to NewExperienceYears")
		}
	})
}

func TestStringExperienceYears(t *testing.T) {
	t.Run("ExperienceYearsが1のとき_「半年未満」がレスポンスされること", func(t *testing.T) {
		experienceYears, err := NewExperienceYears(1)
		if err != nil {
			t.Errorf("failed to newPlanStatus")
		}
		expected := "半年未満"
		actual := experienceYears.String()
		assert.Equal(t, expected, actual)
	})

	t.Run("ExperienceYearsが2のとき_「1年未満」がレスポンスされること", func(t *testing.T) {
		experienceYears, err := NewExperienceYears(2)
		if err != nil {
			t.Errorf("failed to newPlanStatus")
		}
		expected := "1年未満"
		actual := experienceYears.String()
		assert.Equal(t, expected, actual)
	})

	t.Run("ExperienceYearsが3のとき_「3年未満」がレスポンスされること", func(t *testing.T) {
		experienceYears, err := NewExperienceYears(3)
		if err != nil {
			t.Errorf("failed to newPlanStatus")
		}
		expected := "3年未満"
		actual := experienceYears.String()
		assert.Equal(t, expected, actual)
	})

	t.Run("ExperienceYearsが4のとき_「5年未満」がレスポンスされること", func(t *testing.T) {
		experienceYears, err := NewExperienceYears(4)
		if err != nil {
			t.Errorf("failed to newPlanStatus")
		}
		expected := "5年未満"
		actual := experienceYears.String()
		assert.Equal(t, expected, actual)
	})

	t.Run("ExperienceYearsが5のとき_「5年以上」がレスポンスされること", func(t *testing.T) {
		experienceYears, err := NewExperienceYears(5)
		if err != nil {
			t.Errorf("failed to newPlanStatus")
		}
		expected := "5年以上"
		actual := experienceYears.String()
		assert.Equal(t, expected, actual)
	})
}
