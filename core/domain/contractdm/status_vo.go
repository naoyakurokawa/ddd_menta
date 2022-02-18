package contractdm

import "golang.org/x/xerrors"

type Status uint16

const (
	Unapproved Status = iota + 1
	UnderContract
	TerminatedContract
)

func (s Status) Names() []string {
	return []string{
		"未承認",
		"契約中",
		"契約終了",
	}
}

func (s Status) String() (string, error) {
	if s != Unapproved && s != UnderContract && s != TerminatedContract {
		return "", xerrors.New("Status must be 1 or 2 or 3")
	}
	return s.Names()[s-1], nil
}

func (s Status) Uint16() (uint16, error) {
	if s != Unapproved && s != UnderContract && s != TerminatedContract {
		return 0, xerrors.New("Status must be 1 or 2 or 3")
	}
	return uint16(s), nil
}

// コンストラクタ
func NewStatus(s uint16) (Status, error) {
	if s != 1 && s != 2 && s != 3 {
		return 0, xerrors.New("PlanStatus must be 1 or 2 or 3")
	}
	return Status(s), nil
}
