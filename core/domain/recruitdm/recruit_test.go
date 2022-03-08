package recruitdm

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRecruit(t *testing.T) {
	for _, td := range []struct {
		title       string
		inputTitle  string
		budget      uint16
		detail      string
		expectedErr error
	}{
		{
			title:       "titleが空の時_エラーが発生すること",
			inputTitle:  "",
			budget:      rp.budget,
			detail:      rp.detail,
			expectedErr: errors.New("title must not be empty"),
		},
		{
			title:       "titleが255文字を超える場合_エラーが発生すること",
			inputTitle:  strings.Repeat("a", 256),
			budget:      rp.budget,
			detail:      rp.detail,
			expectedErr: errors.New("title must less than 255"),
		},
		{
			title:       "budgetが1000円未満の場合_エラーが発生すること",
			inputTitle:  rp.title,
			budget:      999,
			detail:      rp.detail,
			expectedErr: errors.New("budget more than ¥1000"),
		},
		{
			title:       "detailが空の時_エラーが発生すること",
			inputTitle:  rp.title,
			budget:      rp.budget,
			detail:      "",
			expectedErr: errors.New("title must not be empty"),
		},
		{
			title:       "detailが2000文字を超える場合_エラーが発生すること",
			inputTitle:  rp.title,
			budget:      rp.budget,
			detail:      strings.Repeat("a", 2001),
			expectedErr: errors.New("title must less than 2000"),
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			_, err := NewRecruit(
				rp.recruitID,
				rp.userID,
				td.inputTitle,
				td.budget,
				rp.recruitTypeOnce,
				td.detail,
				rp.recruitStatusDraft,
			)
			if td.expectedErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
