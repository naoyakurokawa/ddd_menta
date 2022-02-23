package contractdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStatus(t *testing.T) {
	t.Run("statusが1のとき_エラーが発生しないこと", func(t *testing.T) {
		_, err := NewStatus(1)
		if err != nil {
			t.Errorf("failed to NewStatus")
		}
	})

	t.Run("ExperienceYearsが4のとき_エラーが発生すること", func(t *testing.T) {
		_, err := NewStatus(4)
		if err == nil {
			t.Errorf("failed to NewStatus")
		}
	})
}

func TestStringStatus(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title  string
		input  uint16
		output string
	}{
		{
			title:  "Statusが1のとき_「未承認」がレスポンスされること",
			input:  uint16(1),
			output: "未承認",
		},
		{
			title:  "Statusが2のとき_「契約中」がレスポンスされること",
			input:  uint16(2),
			output: "契約中",
		},
		{
			title:  "Statusが3のとき_「契約終了」がレスポンスされること",
			input:  uint16(3),
			output: "契約終了",
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			status, err := NewStatus(td.input)
			if err != nil {
				t.Errorf("failed to newPlanStatus")
			}
			actual, err := status.String()
			require.NoError(t, err)
			asserts.Equal(td.output, actual)
		})
	}
}
