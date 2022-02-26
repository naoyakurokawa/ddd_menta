package datamodel

import (
	"time"
)

type Contract struct {
	ContractID     string
	UserID         string
	MentorID       string
	PlanID         string
	ContractStatus uint16
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
