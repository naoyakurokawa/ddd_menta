package suggestiondm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type SuggestionID sharedvo.ID

func NewSuggestionID() SuggestionID {
	return SuggestionID(sharedvo.NewID())
}

func NewSuggestionIDByVal(strId string) (SuggestionID, error) {
	id, err := sharedvo.NewIDByVal(strId)
	if err != nil {
		return SuggestionID(""), xerrors.New("error NewSuggestionIDByVal")
	}
	return SuggestionID(id), nil
}

func NewEmptySuggestionID() SuggestionID {
	return SuggestionID(sharedvo.NewEmptyID())
}

func (i SuggestionID) Equals(i2 SuggestionID) bool {
	return i == i2
}

func (i SuggestionID) String() string {
	return string(i)
}
