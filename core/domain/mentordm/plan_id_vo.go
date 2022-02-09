package mentordm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
)

type PlanID sharedvo.ID

func NewPlanID() PlanID {
	return PlanID(sharedvo.NewID())
}

func NewPlanIDByVal(id string) PlanID {
	return PlanID(id)
}

func NewEmptyPlanID() PlanID {
	return PlanID("")
}
