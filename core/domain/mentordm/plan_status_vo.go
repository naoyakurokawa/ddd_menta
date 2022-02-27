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

func (p PlanStatus) String() string {
	return p.Names()[p-1]
}

func (p PlanStatus) Uint16() uint16 {
	return uint16(p)
}

// コンストラクタ
func NewPlanStatus(planStatus uint16) (PlanStatus, error) {
	if planStatus != 1 && planStatus != 2 && planStatus != 3 {
		return 0, xerrors.New("PlanStatus must be 1 or 2 or 3")
	}
	return PlanStatus(planStatus), nil
}
