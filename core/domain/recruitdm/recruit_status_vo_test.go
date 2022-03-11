package recruitdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func TestNewRecruitStatus(t *testing.T) {
	asserts := assert.New(t)
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
			expectedErr: xerrors.New("RecruitStatus must be 1 or 2 or 3"),
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			_, err := NewRecruitStatus(td.input)
			if td.expectedErr != nil {
				require.Error(t, err)
				asserts.Equal(err.Error(), td.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestStringRecruitStatus(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title       string
		input       uint16
		expected    string
		expectedErr error
	}{
		{
			title:       "RecruitStatusが1のとき_「下書き」がレスポンスされること",
			input:       uint16(1),
			expected:    "下書き",
			expectedErr: nil,
		},
		{
			title:       "RecruitStatusが2のとき_「公開」がレスポンスされること",
			input:       uint16(2),
			expected:    "公開",
			expectedErr: nil,
		},
		{
			title:       "RecruitStatusが3のとき_「募集終了」がレスポンスされること",
			input:       uint16(3),
			expected:    "募集終了",
			expectedErr: nil,
		},
		{
			title:       "RecruitStatusが0のとき_エラーが発生することこと",
			input:       uint16(0),
			expectedErr: xerrors.New("RecruitStatus must be 1 or 2 or 3"),
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			recruitStatus, err := NewRecruitStatus(td.input)
			if td.expectedErr != nil {
				require.Error(t, err)
				asserts.Equal(err.Error(), td.expectedErr.Error())
			} else {
				require.NoError(t, err)
				actual := recruitStatus.String()
				asserts.Equal(actual, td.expected)
			}
		})
	}
}
