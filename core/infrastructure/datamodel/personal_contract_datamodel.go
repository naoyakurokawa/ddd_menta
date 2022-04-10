package datamodel

import (
	"time"
)

type PersonalContract struct {
	PersonalContractID     string
	SuggestionID           string
	PersonalContractStatus uint16
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
