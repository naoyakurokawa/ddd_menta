package recruitdm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRecruitType(t *testing.T) {
	for _, td := range []struct {
		title       string
		input       uint16
		expectedErr error
	}{
		{
			title:       "RecruitTypeが1のとき_エラーが発生しないこと",
			input:       uint16(1),
			expectedErr: nil,
		},
		{
			title:       "RecruitTypeが3のとき_エラーが発生すること",
			input:       uint16(3),
			expectedErr: errors.New("RecruitType must be 1 or 2"),
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			_, err := NewRecruitType(td.input)
			if td.expectedErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestStringRecruitType(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title    string
		input    uint16
		expected string
	}{
		{
			title:    "RecruitStatusが1のとき_「単発」がレスポンスされること",
			input:    uint16(1),
			expected: "単発",
		},
		{
			title:    "RecruitStatusが2のとき_「月額」がレスポンスされること",
			input:    uint16(2),
			expected: "月額",
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			recruitType, err := NewRecruitType(td.input)
			require.NoError(t, err)
			actual := recruitType.String()

			asserts.Equal(actual, td.expected)
		})
	}
}
