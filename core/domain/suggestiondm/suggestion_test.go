package suggestiondm

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSuggestion(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title       string
		price       uint32
		detail      string
		expectedErr error
	}{
		{
			title:       "priceが1000円未満の場合_エラーが発生すること",
			price:       999,
			detail:      sp.detail,
			expectedErr: errors.New("price more than ¥1000"),
		},
		{
			title:       "detailが空の時_エラーが発生すること",
			price:       sp.price,
			detail:      "",
			expectedErr: errors.New("detail must not be empty"),
		},
		{
			title:       "detailが2000文字を超える場合_エラーが発生すること",
			price:       sp.price,
			detail:      strings.Repeat("a", 2001),
			expectedErr: errors.New("detail must less than 2000"),
		},
		{
			title:       "想定通りのSuggestionオブジェクトが生成されること",
			price:       sp.price,
			detail:      sp.detail,
			expectedErr: nil,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			suggestion, err := NewSuggestion(
				sp.suggestionID,
				sp.mentorID,
				sp.recruitID,
				td.price,
				sp.suggestionTypeOnce,
				td.detail,
				sp.suggestionStatusApproval,
			)
			if td.expectedErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				asserts.Equal(suggestion.suggestionID, sp.suggestionID)
				asserts.Equal(suggestion.mentorID, sp.mentorID)
				asserts.Equal(suggestion.recruitID, sp.recruitID)
				asserts.Equal(suggestion.price, td.price)
				asserts.Equal(suggestion.suggestionType, sp.suggestionTypeOnce)
				asserts.Equal(suggestion.detail, sp.detail)
				asserts.Equal(suggestion.suggestionStatus, sp.suggestionStatusApproval)
			}
		})
	}
}

func TestReconstruct(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title            string
		suggestionID     string
		mentorID         string
		recruitID        string
		price            uint32
		suggestionType   uint16
		detail           string
		suggestionStatus uint16
		expectedErr      error
	}{
		{
			title:            "priceが1000円未満の場合_エラーが発生すること",
			suggestionID:     sp.suggestionID.String(),
			recruitID:        sp.recruitID.String(),
			mentorID:         sp.mentorID.String(),
			price:            999,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusApproval.Uint16(),
			expectedErr:      errors.New("price more than ¥1000"),
		},
		{
			title:            "detailが空の時_エラーが発生すること",
			suggestionID:     sp.suggestionID.String(),
			recruitID:        sp.recruitID.String(),
			mentorID:         sp.mentorID.String(),
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           "",
			suggestionStatus: sp.suggestionStatusApproval.Uint16(),
			expectedErr:      errors.New("price more than ¥1000"),
		},
		{
			title:            "detailが2000文字を超える場合_エラーが発生すること",
			suggestionID:     sp.suggestionID.String(),
			recruitID:        sp.recruitID.String(),
			mentorID:         sp.mentorID.String(),
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           strings.Repeat("a", 2001),
			suggestionStatus: sp.suggestionStatusApproval.Uint16(),
			expectedErr:      errors.New("detail must less than 2000"),
		},
		{
			title:            "suggestionIDが空の場合_エラーが発生すること",
			suggestionID:     "",
			recruitID:        sp.recruitID.String(),
			mentorID:         sp.mentorID.String(),
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusApproval.Uint16(),
			expectedErr:      errors.New("error NewRecruitIDByVal"),
		},
		{
			title:            "recruitIDが空の場合_エラーが発生すること",
			suggestionID:     sp.suggestionID.String(),
			recruitID:        "",
			mentorID:         sp.mentorID.String(),
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusApproval.Uint16(),
			expectedErr:      errors.New("error NewRecruitIDByVal"),
		},
		{
			title:            "mentorIDが空の場合_エラーが発生すること",
			suggestionID:     sp.suggestionID.String(),
			recruitID:        sp.recruitID.String(),
			mentorID:         "",
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusApproval.Uint16(),
			expectedErr:      errors.New("error NewMentorIDByVal"),
		},
		{
			title:            "suggestionTypeが0の場合_エラーが発生すること",
			suggestionID:     sp.suggestionID.String(),
			recruitID:        sp.recruitID.String(),
			mentorID:         sp.mentorID.String(),
			price:            sp.price,
			suggestionType:   0,
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusApproval.Uint16(),
			expectedErr:      errors.New("error NewSuggestionType"),
		},
		{
			title:            "suggestionStatusが0の場合_エラーが発生すること",
			suggestionID:     sp.suggestionID.String(),
			recruitID:        sp.recruitID.String(),
			mentorID:         sp.mentorID.String(),
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: 0,
			expectedErr:      errors.New("error NewSuggestionStatus"),
		},
		{
			title:            "想定通りのSuggestionオブジェクトが生成されること",
			suggestionID:     sp.suggestionID.String(),
			recruitID:        sp.recruitID.String(),
			mentorID:         sp.mentorID.String(),
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusApproval.Uint16(),
			expectedErr:      nil,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			suggestion, err := Reconstruct(
				td.suggestionID,
				td.mentorID,
				td.recruitID,
				td.price,
				td.suggestionType,
				td.detail,
				td.suggestionStatus,
			)
			if td.expectedErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				asserts.Equal(suggestion.suggestionID, sp.suggestionID)
				asserts.Equal(suggestion.mentorID, sp.mentorID)
				asserts.Equal(suggestion.recruitID, sp.recruitID)
				asserts.Equal(suggestion.price, sp.price)
				asserts.Equal(suggestion.suggestionType, sp.suggestionTypeOnce)
				asserts.Equal(suggestion.detail, sp.detail)
				asserts.Equal(suggestion.suggestionStatus, sp.suggestionStatusApproval)
			}
		})
	}
}
