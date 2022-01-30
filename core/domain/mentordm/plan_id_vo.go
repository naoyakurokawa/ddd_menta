package mentordm

import (
	"github.com/google/uuid"
)

type PlanID string

func newPlanID() (PlanID, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return PlanID(""), err
	}
	us := u.String()

	return PlanID(us), nil
}

func (m PlanID) Equals(m2 PlanID) bool {
	return m.Value() == m2.Value()
}

func (m PlanID) Value() string {
	return string(m)
}

func PlanIDType(strPlanID string) PlanID {
	return PlanID(strPlanID)
}
