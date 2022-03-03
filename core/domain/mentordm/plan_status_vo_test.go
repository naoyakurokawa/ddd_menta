package mentordm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlanStatus(t *testing.T) {
	t.Run("PlanStatusが1のとき_エラーが発生しないこと", func(t *testing.T) {
		_, err := NewPlanStatus(1)
		if err != nil {
			t.Errorf("failed to NewPlanStatus")
		}
	})

	t.Run("PlanStatusが4のとき_エラーが発生すること", func(t *testing.T) {
		_, err := NewPlanStatus(4)
		if err == nil {
			t.Errorf("failed to NewPlanStatus")
		}
	})
}
func TestString(t *testing.T) {
	t.Run("PlanStatusが1のとき_「相談できます」がレスポンスされること", func(t *testing.T) {
		planStatus1, err := NewPlanStatus(1)
		if err != nil {
			t.Errorf("failed to newPlanStatus")
		}
		expected := "相談できます"
		actual := planStatus1.String()
		assert.Equal(t, expected, actual)
	})

	t.Run("PlanStatusが2のとき_「今、忙しいです」がレスポンスされること", func(t *testing.T) {
		planStatus2, err := NewPlanStatus(2)
		if err != nil {
			t.Errorf("failed to newPlanStatus")
		}
		expected := "今、忙しいです"
		actual := planStatus2.String()
		assert.Equal(t, expected, actual)
	})

	t.Run("PlanStatusが3のとき_「表示しない」がレスポンスされること", func(t *testing.T) {
		planStatus3, err := NewPlanStatus(3)
		if err != nil {
			t.Errorf("failed to newPlanStatus")
		}
		expected := "表示しない"
		actual := planStatus3.String()
		assert.Equal(t, expected, actual)
	})

}

func TestUint16(t *testing.T) {
	t.Run("PlanStatusが1のとき_1がレスポンスされること", func(t *testing.T) {
		planStatus1, err := NewPlanStatus(1)
		if err != nil {
			t.Errorf("failed to newPlanStatus")
		}
		expected := uint16(1)
		actual := planStatus1.Uint16()
		assert.Equal(t, expected, actual)
	})
}
