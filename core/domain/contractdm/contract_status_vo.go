package contractdm

import "golang.org/x/xerrors"

type ContractStatus uint16

const (
	Unapproved ContractStatus = iota + 1
	UnderContract
	TerminatedContract
)

func (s ContractStatus) Names() []string {
	return []string{
		"未承認",
		"契約中",
		"契約終了",
	}
}

func (s ContractStatus) String() (string, error) {
	if !s.isWithinRange() {
		return "", xerrors.New("ContractStatus must be 1 or 2 or 3")
	}
	return s.Names()[s-1], nil
}

func (s ContractStatus) Uint16() uint16 {
	return uint16(s)
}

func NewContractStatus(s uint16) (ContractStatus, error) {
	if s != 1 && s != 2 && s != 3 {
		return 0, xerrors.New("ContractStatus must be 1 or 2 or 3")
	}
	return ContractStatus(s), nil
}

func (s ContractStatus) isWithinRange() bool {
	return s == Unapproved || s == UnderContract || s == TerminatedContract
}
