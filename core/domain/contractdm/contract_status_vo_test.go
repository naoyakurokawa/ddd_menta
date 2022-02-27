package contractdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewContractStatus(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title  string
		input  uint16
		output string
	}{
		{
			title:  "ContractStatusが1のとき_エラーが発生しないこと",
			input:  uint16(1),
			output: "",
		},
		{
			title:  "ExperienceYearsが4のとき_エラーが発生すること",
			input:  uint16(4),
			output: "ContractStatus must be 1 or 2 or 3",
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			_, err := NewContractStatus(td.input)
			strErr := ""
			if err != nil {
				strErr = err.Error()
			}
			asserts.Equal(td.output, strErr)
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
			title:  "ContractStatusが1のとき_「未承認」がレスポンスされること",
			input:  uint16(1),
			output: "未承認",
		},
		{
			title:  "ContractStatusが2のとき_「契約中」がレスポンスされること",
			input:  uint16(2),
			output: "契約中",
		},
		{
			title:  "ContractStatusが3のとき_「契約終了」がレスポンスされること",
			input:  uint16(3),
			output: "契約終了",
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			contractStatus, err := NewContractStatus(td.input)
			if err != nil {
				t.Errorf("failed to NewContractStatus")
			}
			actual, err := contractStatus.String()
			require.NoError(t, err)
			asserts.Equal(td.output, actual)
		})
	}
}
