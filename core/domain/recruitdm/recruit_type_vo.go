package recruitdm

import "golang.org/x/xerrors"

type RecruitType uint16

const (
	Once RecruitType = iota + 1
	Subscription
)

func (r RecruitType) Names() []string {
	return []string{
		"単発",
		"月額",
	}
}

func (r RecruitType) String() string {
	return r.Names()[r-1]
}

func (r RecruitType) Uint16() uint16 {
	return uint16(r)
}

func NewRecruitType(recruitType uint16) (RecruitType, error) {
	if recruitType != 1 && 2 != recruitType {
		return 0, xerrors.New("RecruitType must be 1 or 2")
	}
	return RecruitType(recruitType), nil
}
