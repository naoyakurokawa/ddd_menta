package datamodel

import (
	"time"
)

type Recruit struct {
	RecruitID     string
	UserID        string
	Title         string
	Budget        uint32
	RecruitType   uint16
	Detail        string
	RecruitStatus uint16
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
