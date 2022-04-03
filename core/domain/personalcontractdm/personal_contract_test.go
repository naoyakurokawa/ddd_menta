package personalcontractdm

import (
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func TestNewContract(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title                  string
		personalContractID     PersonalContractID
		suggestionID           suggestiondm.SuggestionID
		personalContractStatus PersonalContractStatus
	}{
		{
			title:                  "正しいパラメータの場合_エラーが発生せず_想定通りのPersonalContractが生成されること",
			personalContractID:     pp.personalContractID,
			suggestionID:           pp.suggestionID,
			personalContractStatus: pp.personalContractStatus,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			actual, err := NewPersonalContract(
				pp.personalContractID,
				pp.suggestionID,
				pp.personalContractStatus,
			)
			require.NoError(t, err)

			expected := &PersonalContract{
				personalContractID:     pp.personalContractID,
				suggestionID:           pp.suggestionID,
				personalContractStatus: pp.personalContractStatus,
				createdAt:              actual.CreatedAt(),
				updatedAt:              actual.UpdatedAt(),
			}
			asserts.Equal(actual, expected)
		})
	}
}

func TestReconstruct(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title                  string
		personalContractID     string
		suggestionID           string
		personalContractStatus uint16
		expectedErr            error
	}{
		{
			title:                  "正しいパラメータの場合_エラーが発生せず_想定通りのPersonalContractが生成されること",
			personalContractID:     pp.personalContractID.String(),
			suggestionID:           pp.suggestionID.String(),
			personalContractStatus: pp.personalContractStatus.Uint16(),
			expectedErr:            nil,
		},
		{
			title:                  "personalContractIDが空文字の時_エラーが発生すること",
			personalContractID:     "",
			suggestionID:           pp.suggestionID.String(),
			personalContractStatus: pp.personalContractStatus.Uint16(),
			expectedErr:            xerrors.New("error NewPersonalContractIDByVal"),
		},
		{
			title:                  "suggestionIDが空文字の時_エラーが発生すること",
			personalContractID:     pp.personalContractID.String(),
			suggestionID:           "",
			personalContractStatus: pp.personalContractStatus.Uint16(),
			expectedErr:            xerrors.New("error NewSuggestionIDByVal"),
		},
		{
			title:                  "personalContractStatusが0の時_エラーが発生すること",
			personalContractID:     pp.personalContractID.String(),
			suggestionID:           pp.suggestionID.String(),
			personalContractStatus: 0,
			expectedErr:            xerrors.New("error NewPersonalContractStatus"),
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			actual, err := Reconstruct(
				td.personalContractID,
				td.suggestionID,
				td.personalContractStatus,
			)

			if td.expectedErr == nil {
				require.NoError(t, err)
				expected := &PersonalContract{
					personalContractID:     pp.personalContractID,
					suggestionID:           pp.suggestionID,
					personalContractStatus: pp.personalContractStatus,
					createdAt:              actual.CreatedAt(),
					updatedAt:              actual.UpdatedAt(),
				}
				asserts.Equal(actual, expected)
			} else {
				asserts.Equal(td.expectedErr.Error(), err.Error())
			}
		})
	}
}
