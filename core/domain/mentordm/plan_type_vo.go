package mentordm

import "golang.org/x/xerrors"

type PlanType uint16

const (
	Once PlanType = iota
	Subscription
)

func (p PlanType) Names() []string {
	return []string{
		"単発",
		"月額",
	}
}

func (p PlanType) String() (string, error) {
	if p != Once || Subscription != p {
		return "", xerrors.New("PlanType must be 0 or 1")
	}
	return p.Names()[p], nil
}

func (p PlanType) Uint16() (uint16, error) {
	if p != Once || Subscription != p {
		return 0, xerrors.New("PlanType must be 0 or 1")
	}
	return uint16(p), nil
}

// コンストラクタ
func newPlanType(planType uint16) (PlanType, error) {
	if planType != 0 || 1 != planType {
		return 0, xerrors.New("PlanType must be 0 or 1")
	}
	return PlanType(planType), nil
}
