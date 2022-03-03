package mentordm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlanType(t *testing.T) {
	t.Run("PlanTypeが1のとき_エラーが発生しないこと", func(t *testing.T) {
		_, err := NewPlanType(1)
		if err != nil {
			t.Errorf("failed to NewPlanType")
		}
	})

	t.Run("PlanTypeが3のとき_エラーが発生すること", func(t *testing.T) {
		_, err := NewPlanType(3)
		if err == nil {
			t.Errorf("failed to NewPlanType")
		}
	})
}

func TestStringPlanType(t *testing.T) {
	t.Run("PlanTypeが1のとき_「単発」がレスポンスされること", func(t *testing.T) {
		planType1, err := NewPlanType(1)
		if err != nil {
			t.Errorf("failed to NewPlanType")
		}
		expected := "単発"
		actual := planType1.String()
		assert.Equal(t, expected, actual)
	})

	t.Run("PlanTypeが2のとき_「月額」がレスポンスされること", func(t *testing.T) {
		planType2, err := NewPlanType(2)
		if err != nil {
			t.Errorf("failed to NewPlanType")
		}
		expected := "月額"
		actual := planType2.String()
		assert.Equal(t, expected, actual)
	})
}

func TestUint16PlanType(t *testing.T) {
	t.Run("PlanTypeが2のとき_2がレスポンスされること", func(t *testing.T) {
		planType2, err := NewPlanType(2)
		if err != nil {
			t.Errorf("failed to NewPlanType")
		}
		expected := uint16(2)
		actual := planType2.Uint16()
		assert.Equal(t, expected, actual)
	})
}
