package personalcontractdm

import "golang.org/x/xerrors"

type PersonalContractStatus uint16

const (
	Unapproved PersonalContractStatus = iota + 1
	UnderContract
	TerminatedContract
)

func (s PersonalContractStatus) Names() []string {
	return []string{
		"未承認",
		"契約中",
		"契約終了",
	}
}

func (s PersonalContractStatus) String() (string, error) {
	if !s.isWithinRange() {
		return "", xerrors.New("PersonalContractStatus must be 1 or 2 or 3")
	}
	return s.Names()[s-1], nil
}

func (s PersonalContractStatus) Uint16() uint16 {
	return uint16(s)
}

func NewPersonalContractStatus(s uint16) (PersonalContractStatus, error) {
	if s != 1 && s != 2 && s != 3 {
		return 0, xerrors.New("PersonalContractStatus must be 1 or 2 or 3")
	}
	return PersonalContractStatus(s), nil
}

func (s PersonalContractStatus) isWithinRange() bool {
	return s == Unapproved || s == UnderContract || s == TerminatedContract
}
