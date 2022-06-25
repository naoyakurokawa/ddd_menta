package suggestiondm

import "golang.org/x/xerrors"

type SuggestionStatus uint16

const (
	Unapproved SuggestionStatus = iota + 1
	Approval
	Terminated
)

func (s SuggestionStatus) Names() []string {
	return []string{
		"提案中",
		"承認",
		"終了",
	}
}

func (s SuggestionStatus) String() string {
	return s.Names()[s-1]
}

func (s SuggestionStatus) Uint16() uint16 {
	return uint16(s)
}

func NewSuggestionStatus(s uint16) (SuggestionStatus, error) {
	if s != 1 && s != 2 && s != 3 {
		return 0, xerrors.New("SuggestionStatus must be 1 or 2 or 3")
	}
	return SuggestionStatus(s), nil
}
