package datamodel

import (
	"time"
)

type Suggestion struct {
	SuggestionID     string
	MentorID         string
	RecruitID        string
	Price            uint32
	SuggestionType   uint16
	Detail           string
	SuggestionStatus uint16
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
