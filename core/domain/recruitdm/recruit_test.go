package recruitdm

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRecruit(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title       string
		inputTitle  string
		budget      uint32
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
		{
			title:       "想定通りのRecruitオブジェクトが生成されること",
			inputTitle:  rp.title,
			budget:      rp.budget,
			detail:      rp.detail,
			expectedErr: nil,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			recruit, err := NewRecruit(
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
				asserts.Equal(recruit.recruitID, rp.recruitID)
				asserts.Equal(recruit.userID, rp.userID)
				asserts.Equal(recruit.title, rp.title)
				asserts.Equal(recruit.budget, rp.budget)
				asserts.Equal(recruit.recruitType, rp.recruitTypeOnce)
				asserts.Equal(recruit.detail, rp.detail)
				asserts.Equal(recruit.recruitStatus, rp.recruitStatusDraft)
			}
		})
	}
}

func TestReconstruct(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title         string
		recruitID     string
		userID        string
		inputTitle    string
		budget        uint32
		recruitType   uint16
		detail        string
		recruitStatus uint16
		expectedErr   error
	}{
		{
			title:         "titleが空の時_エラーが発生すること",
			recruitID:     rp.recruitID.String(),
			userID:        rp.userID.String(),
			inputTitle:    "",
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   errors.New("title must not be empty"),
		},
		{
			title:         "titleが255文字を超える場合_エラーが発生すること",
			recruitID:     rp.recruitID.String(),
			userID:        rp.userID.String(),
			inputTitle:    strings.Repeat("a", 256),
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   errors.New("title must less than 255"),
		},
		{
			title:         "budgetが1000円未満の場合_エラーが発生すること",
			recruitID:     rp.recruitID.String(),
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        999,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   errors.New("budget more than ¥1000"),
		},
		{
			title:         "detailが空の時_エラーが発生すること",
			recruitID:     rp.recruitID.String(),
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        "",
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   errors.New("title must not be empty"),
		},
		{
			title:         "detailが2000文字を超える場合_エラーが発生すること",
			recruitID:     rp.recruitID.String(),
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        strings.Repeat("a", 2001),
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   errors.New("title must less than 2000"),
		},
		{
			title:         "recruitIDが空場合_エラーが発生すること",
			recruitID:     "",
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   errors.New("error NewRecruitIDByVal"),
		},
		{
			title:         "userIDが空場合_エラーが発生すること",
			recruitID:     rp.recruitID.String(),
			userID:        "",
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   errors.New("error NewUserIDByVal"),
		},
		{
			title:         "recruitTypeが0場合_エラーが発生すること",
			recruitID:     rp.recruitID.String(),
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   0,
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   errors.New("error NewRecruitType"),
		},
		{
			title:         "recruitStatusが0場合_エラーが発生すること",
			recruitID:     rp.recruitID.String(),
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: 0,
			expectedErr:   errors.New("error NewRecruitStatus"),
		},
		{
			title:         "想定通りのRecruitオブジェクトが生成されること",
			recruitID:     rp.recruitID.String(),
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   nil,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			recruit, err := Reconstruct(
				td.recruitID,
				td.userID,
				td.inputTitle,
				td.budget,
				td.recruitType,
				td.detail,
				td.recruitStatus,
			)
			if td.expectedErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				asserts.Equal(recruit.recruitID, rp.recruitID)
				asserts.Equal(recruit.userID, rp.userID)
				asserts.Equal(recruit.title, rp.title)
				asserts.Equal(recruit.budget, rp.budget)
				asserts.Equal(recruit.recruitType, rp.recruitTypeOnce)
				asserts.Equal(recruit.detail, rp.detail)
				asserts.Equal(recruit.recruitStatus, rp.recruitStatusDraft)
			}
		})
	}
}
