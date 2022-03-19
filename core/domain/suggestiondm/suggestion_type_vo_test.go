package suggestiondm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func TestNewSuggestionType(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title       string
		input       uint16
		expectedErr error
	}{
		{
			title:       "SuggestionTypeが1のとき_エラーが発生しないこと",
			input:       uint16(1),
			expectedErr: nil,
		},
		{
			title:       "SuggestionTypeが3のとき_エラーが発生すること",
			input:       uint16(3),
			expectedErr: xerrors.New("SuggestionType must be 1 or 2"),
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			_, err := NewSuggestionType(td.input)
			if td.expectedErr != nil {
				require.Error(t, err)
				asserts.Equal(err.Error(), td.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestStringSuggestionType(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title    string
		input    uint16
		expected string
	}{
		{
			title:    "SuggestionStatusが1のとき_「単発」がレスポンスされること",
			input:    uint16(1),
			expected: "単発",
		},
		{
			title:    "SuggestionStatusが2のとき_「月額」がレスポンスされること",
			input:    uint16(2),
			expected: "月額",
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			suggestionType, err := NewSuggestionType(td.input)
			require.NoError(t, err)
			actual := suggestionType.String()

			asserts.Equal(actual, td.expected)
		})
	}
}
