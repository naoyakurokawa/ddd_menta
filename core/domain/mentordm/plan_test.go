package mentordm

import (
	"strings"
	"testing"
)

func TestNewPlan(t *testing.T) {
	const (
		title    = "DDDのメンタリング"
		category = "設計"
		tag      = "DDD"
		detial   = "DDDの設計手法を学べます"
		price    = 10000
	)
	planID := NewPlanID()
	planType1, err := NewPlanType(1)
	if err != nil {
		t.Errorf("failed to NewPlanType")
	}

	planStatus1, err := NewPlanStatus(1)
	if err != nil {
		t.Errorf("failed to newPlanStatus")
	}

	t.Run("titleが空の場合_エラーとなること", func(t *testing.T) {
		blankTitle := ""
		_, err := NewPlan(planID, blankTitle, category, tag, detial, planType1, price, planStatus1)
		if err == nil {
			t.Errorf("failed to title blank validation")
		}
	})

	t.Run("titleが255文字を超えるの場合_エラーとなること", func(t *testing.T) {
		overTitle := strings.Repeat("a", 256)
		_, err := NewPlan(planID, overTitle, category, tag, detial, planType1, price, planStatus1)
		if err == nil {
			t.Errorf("failed to title maxlength validation: %v", overTitle)
		}
	})

	t.Run("categoryが空の場合_エラーとなること", func(t *testing.T) {
		blankCategory := ""
		_, err := NewPlan(planID, title, blankCategory, tag, detial, planType1, price, planStatus1)
		if err == nil {
			t.Errorf("failed to title blank validation")
		}
	})

	t.Run("tagが空の場合_エラーとなること", func(t *testing.T) {
		blankTag := ""
		_, err := NewPlan(planID, title, category, blankTag, detial, planType1, price, planStatus1)
		if err == nil {
			t.Errorf("failed to tag blank validation")
		}
	})

	t.Run("detialが空の場合_エラーとなること", func(t *testing.T) {
		blankDetial := ""
		_, err := NewPlan(planID, title, category, tag, blankDetial, planType1, price, planStatus1)
		if err == nil {
			t.Errorf("failed to detial blank validation")
		}
	})

	t.Run("detialが2000文字を超えるの場合_エラーとなること", func(t *testing.T) {
		overDetail := strings.Repeat("a", 2001)
		_, err := NewPlan(planID, title, category, tag, overDetail, planType1, price, planStatus1)
		if err == nil {
			t.Errorf("failed to detail maxlength validation: %v", overDetail)
		}
	})

	t.Run("priceが0の場合_エラーとなること", func(t *testing.T) {
		errPrice := uint16(0)
		_, err := NewPlan(planID, title, category, tag, detial, planType1, errPrice, planStatus1)
		if err == nil {
			t.Errorf("failed to price 0 validation")
		}
	})
}
