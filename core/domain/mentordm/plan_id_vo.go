package mentordm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type PlanID sharedvo.ID

func NewPlanID() PlanID {
	return PlanID(sharedvo.NewID())
}

func NewPlanIDByVal(id string) (PlanID, error) {
	ID, err := sharedvo.NewIDByVal(id)
	if err != nil {
		return PlanID(""), xerrors.New("error NewMentorIDByVal")
	}
	return PlanID(ID), nil
}

func NewEmptyPlanID() PlanID {
	return PlanID(sharedvo.NewEmptyID())
}
