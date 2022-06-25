package recruitdm

import "golang.org/x/xerrors"

type RecruitStatus uint16

const (
	Draft RecruitStatus = iota + 1
	Published
	Terminated
)

func (s RecruitStatus) Names() []string {
	return []string{
		"下書き",
		"公開",
		"募集終了",
	}
}

func (s RecruitStatus) String() string {
	return s.Names()[s-1]
}

func (s RecruitStatus) Uint16() uint16 {
	return uint16(s)
}

func NewRecruitStatus(s uint16) (RecruitStatus, error) {
	if s != 1 && s != 2 && s != 3 {
		return 0, xerrors.New("RecruitStatus must be 1 or 2 or 3")
	}
	return RecruitStatus(s), nil
}
