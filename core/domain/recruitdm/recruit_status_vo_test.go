package recruitdm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRecruitStatus(t *testing.T) {
	for _, td := range []struct {
		title       string
		input       uint16
		expectedErr error
	}{
		{
			title:       "RecruitStatusが1のとき_エラーが発生しないこと",
			input:       uint16(1),
			expectedErr: nil,
		},
		{
			title:       "RecruitStatusが4のとき_エラーが発生すること",
			input:       uint16(4),
			expectedErr: errors.New("RecruitStatus must be 1 or 2 or 3"),
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			_, err := NewRecruitStatus(td.input)
			if td.expectedErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestStringRecruitStatus(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title    string
		input    uint16
		expected string
	}{
		{
			title:    "RecruitStatusが1のとき_「下書き」がレスポンスされること",
			input:    uint16(1),
			expected: "下書き",
		},
		{
			title:    "RecruitStatusが2のとき_「公開」がレスポンスされること",
			input:    uint16(2),
			expected: "公開",
		},
		{
			title:    "RecruitStatusが3のとき_「募集終了」がレスポンスされること",
			input:    uint16(3),
			expected: "募集終了",
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			recruitStatus, err := NewRecruitStatus(td.input)
			require.NoError(t, err)
			actual := recruitStatus.String()

			asserts.Equal(actual, td.expected)
		})
	}
}
