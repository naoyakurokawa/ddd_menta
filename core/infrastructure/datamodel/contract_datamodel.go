package datamodel

import (
	"time"
)

type Contract struct {
	ContractID string
	UserID     string
	PlanID     string
	Status     uint16
	CreatedAt  time.Time
}
