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
	asserts := assert.New(t)
	for _, td := range []struct {
		title  string
		input  uint16
		output string
	}{
		{
			title:  "ExperienceYearsが1のとき_「半年未満」がレスポンスされること",
			input:  uint16(1),
			output: "半年未満",
		},
		{
			title:  "ExperienceYearsが2のとき_「1年未満」がレスポンスされること",
			input:  uint16(2),
			output: "1年未満",
		},
		{
			title:  "ExperienceYearsが3のとき_「3年未満」がレスポンスされること",
			input:  uint16(3),
			output: "3年未満",
		},
		{
			title:  "ExperienceYearsが4のとき_「5年未満」がレスポンスされること",
			input:  uint16(4),
			output: "5年未満",
		},
		{
			title:  "ExperienceYearsが5のとき_「5年以上」がレスポンスされること",
			input:  uint16(5),
			output: "5年以上",
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			experienceYears, err := NewExperienceYears(td.input)
			if err != nil {
				t.Errorf("failed to newPlanStatus")
			}
			actual := experienceYears.String()
			asserts.Equal(td.output, actual)
		})
	}
}
