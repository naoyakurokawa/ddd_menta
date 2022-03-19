package suggestiondm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func TestNewSuggestionStatus(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title       string
		input       uint16
		expectedErr error
	}{
		{
			title:       "SuggestionStatusが1のとき_エラーが発生しないこと",
			input:       uint16(1),
			expectedErr: nil,
		},
		{
			title:       "SuggestionStatusが4のとき_エラーが発生すること",
			input:       uint16(4),
			expectedErr: xerrors.New("SuggestionStatus must be 1 or 2 or 3"),
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			_, err := NewSuggestionStatus(td.input)
			if td.expectedErr != nil {
				require.Error(t, err)
				asserts.Equal(err.Error(), td.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestStringSuggestionStatus(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title       string
		input       uint16
		expected    string
		expectedErr error
	}{
		{
			title:       "SuggestionStatusが1のとき_「提案中」がレスポンスされること",
			input:       uint16(1),
			expected:    "提案中",
			expectedErr: nil,
		},
		{
			title:       "SuggestionStatusが2のとき_「承認」がレスポンスされること",
			input:       uint16(2),
			expected:    "承認",
			expectedErr: nil,
		},
		{
			title:       "SuggestionStatusが3のとき_「終了」がレスポンスされること",
			input:       uint16(3),
			expected:    "終了",
			expectedErr: nil,
		},
		{
			title:       "SuggestionStatusが0のとき_エラーが発生することこと",
			input:       uint16(0),
			expectedErr: xerrors.New("SuggestionStatus must be 1 or 2 or 3"),
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			suggestionStatus, err := NewSuggestionStatus(td.input)
			if td.expectedErr != nil {
				require.Error(t, err)
				asserts.Equal(err.Error(), td.expectedErr.Error())
			} else {
				require.NoError(t, err)
				actual := suggestionStatus.String()
				asserts.Equal(actual, td.expected)
			}
		})
	}
}
