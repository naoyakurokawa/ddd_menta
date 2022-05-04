package personalcontractdm

import (
	"testing"

	"github.com/naoyakurokawa/ddd_menta/customerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPersonalContractStatus(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title       string
		input       uint16
		expectedErr error
	}{
		{
			title:       "PersonalContractStatusが1のとき_エラーが発生しないこと",
			input:       uint16(1),
			expectedErr: nil,
		},
		{
			title:       "PersonalContractStatusが4のとき_エラーが発生すること",
			input:       uint16(4),
			expectedErr: customerrors.NewInvalidParameter("PersonalContractStatus must be 1 or 2 or 3"),
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			_, err := NewPersonalContractStatus(td.input)
			if td.expectedErr == nil {
				require.NoError(t, err)
			} else {
				asserts.Equal(td.expectedErr.Error(), err.Error())
			}
		})
	}
}

func TestStringContractStatus(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title  string
		input  uint16
		output string
	}{
		{
			title:  "PersonalContractStatusが1のとき_「未承認」がレスポンスされること",
			input:  uint16(1),
			output: "未承認",
		},
		{
			title:  "PersonalContractStatusが2のとき_「契約中」がレスポンスされること",
			input:  uint16(2),
			output: "契約中",
		},
		{
			title:  "PersonalContractStatusが3のとき_「契約終了」がレスポンスされること",
			input:  uint16(3),
			output: "契約終了",
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			personalContractStatus, err := NewPersonalContractStatus(td.input)
			if err != nil {
				t.Errorf("failed to NewPersonalContractStatus")
			}
			actual, err := personalContractStatus.String()
			require.NoError(t, err)
			asserts.Equal(td.output, actual)
		})
	}
}
