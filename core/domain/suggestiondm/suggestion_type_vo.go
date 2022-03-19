package suggestiondm

import "golang.org/x/xerrors"

type SuggestionType uint16

const (
	Once SuggestionType = iota + 1
	Subscription
)

func (s SuggestionType) Names() []string {
	return []string{
		"単発",
		"月額",
	}
}

func (s SuggestionType) String() string {
	return s.Names()[s-1]
}

func (s SuggestionType) Uint16() uint16 {
	return uint16(s)
}

func NewSuggestionType(suggestionType uint16) (SuggestionType, error) {
	if suggestionType != 1 && 2 != suggestionType {
		return 0, xerrors.New("SuggestionType must be 1 or 2")
	}
	return SuggestionType(suggestionType), nil
}
