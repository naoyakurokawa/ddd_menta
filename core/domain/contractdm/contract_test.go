package contractdm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewContract(t *testing.T) {
	t.Run("実行時に_エラーが発生しないこと", func(t *testing.T) {
		_, err := NewContract(
			cp.contractID,
			cp.userID,
			cp.mentorID,
			cp.planID,
			cp.status,
		)
		require.NoError(t, err)
	})
}
