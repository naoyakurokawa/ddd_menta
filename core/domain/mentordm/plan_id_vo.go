package mentordm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type PlanID sharedvo.ID

func NewPlanID() PlanID {
	return PlanID(sharedvo.NewID())
}

func NewPlanIDByVal(srtID string) (PlanID, error) {
	id, err := sharedvo.NewIDByVal(srtID)
	if err != nil {
		return PlanID(""), xerrors.New("error NewPlanIDByVal")
	}
	return PlanID(id), nil
}

func NewEmptyPlanID() PlanID {
	return PlanID(sharedvo.NewEmptyID())
}

func (i PlanID) String() string {
	return string(i)
}
