package mentordm

import "golang.org/x/xerrors"

type PlanType uint16

const (
	Once PlanType = iota + 1
	Subscription
)

func (p PlanType) Names() []string {
	return []string{
		"単発",
		"月額",
	}
}

func (p PlanType) String() string {
	return p.Names()[p-1]
}

func (p PlanType) Uint16() uint16 {
	return uint16(p)
}

func NewPlanType(planType uint16) (PlanType, error) {
	if planType != 1 && 2 != planType {
		return 0, xerrors.New("PlanType must be 1 or 2")
	}
	return PlanType(planType), nil
}
