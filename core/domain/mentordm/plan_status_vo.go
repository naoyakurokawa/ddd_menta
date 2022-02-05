package mentordm

import "golang.org/x/xerrors"

type PlanStatus uint16

const (
	Active PlanStatus = iota + 1
	Busy
	Close
)

func (p PlanStatus) Names() []string {
	return []string{
		"相談できます",
		"今、忙しいです",
		"表示しない",
	}
}

func (p PlanStatus) String() (string, error) {
	if p != Active || p != Busy || p != Close {
		return "", xerrors.New("PlanStatus must be 0 or 1 or 2")
	}
	return p.Names()[p], nil
}

func (p PlanStatus) Uint16() (uint16, error) {
	if p != Active || p != Busy || p != Close {
		return 0, xerrors.New("PlanStatus must be 0 or 1 or 2")
	}
	return uint16(p), nil
}

// コンストラクタ
func newPlanStatus(planStatus uint16) (PlanStatus, error) {
	if planStatus != 0 || planStatus != 1 || planStatus != 2 {
		return 0, xerrors.New("PlanStatus must be 0 or 1 or 2")
	}
	return PlanStatus(planStatus), nil
}
