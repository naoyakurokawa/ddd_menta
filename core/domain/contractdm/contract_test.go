package contractdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewContract(t *testing.T) {
	t.Run("実行時に_エラーが発生しないこと", func(t *testing.T) {
		_, err := NewContract(
			cp.contractID,
			cp.userID,
			cp.mentorID,
			cp.planID,
			cp.contractStatus,
		)
		require.NoError(t, err)
	})
}

func TestReconstruct(t *testing.T) {
	t.Run("実行時に_エラーが発生しないこと", func(t *testing.T) {
		_, err := Reconstruct(
			cp.contractID.String(),
			cp.userID.String(),
			cp.mentorID.String(),
			cp.planID.String(),
			cp.contractStatus.Uint16(),
		)
		require.NoError(t, err)
	})
}

func TestCanUpdateContractStatus(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title                 string
		contractStatus        uint16
		requestContractStatus uint16
		expected              bool
	}{
		{
			title:                 "ContractStatusを1から1に更新するとき_falseとなること",
			contractStatus:        uint16(1),
			requestContractStatus: uint16(1),
			expected:              false,
		},
		{
			title:                 "ContractStatusを1から2に更新するとき_trueとなること",
			contractStatus:        uint16(1),
			requestContractStatus: uint16(2),
			expected:              true,
		},
		{
			title:                 "ContractStatusを1から3に更新するとき_falseとなること",
			contractStatus:        uint16(1),
			requestContractStatus: uint16(3),
			expected:              false,
		},
		{
			title:                 "ContractStatusを2から1に更新するとき_falseとなること",
			contractStatus:        uint16(2),
			requestContractStatus: uint16(1),
			expected:              false,
		},
		{
			title:                 "ContractStatusを2から2に更新するとき_falseとなること",
			contractStatus:        uint16(2),
			requestContractStatus: uint16(2),
			expected:              false,
		},
		{
			title:                 "ContractStatusを2から3に更新するとき_trueとなること",
			contractStatus:        uint16(2),
			requestContractStatus: uint16(3),
			expected:              true,
		},
		{
			title:                 "ContractStatusを3から1に更新するとき_falseとなること",
			contractStatus:        uint16(3),
			requestContractStatus: uint16(1),
			expected:              false,
		},
		{
			title:                 "ContractStatusを3から2に更新するとき_falseとなること",
			contractStatus:        uint16(3),
			requestContractStatus: uint16(2),
			expected:              false,
		},
		{
			title:                 "ContractStatusを3から3に更新するとき_falseとなること",
			contractStatus:        uint16(3),
			requestContractStatus: uint16(3),
			expected:              false,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			contractStatus, err := NewContractStatus(td.contractStatus)
			require.NoError(t, err)
			contract, err := NewContract(
				cp.contractID,
				cp.userID,
				cp.mentorID,
				cp.planID,
				contractStatus,
			)
			require.NoError(t, err)

			requestContractStatus, err := NewContractStatus(td.requestContractStatus)
			require.NoError(t, err)

			actual := contract.CanUpdateContractStatus(requestContractStatus)
			asserts.Equal(td.expected, actual)
		})
	}
}
