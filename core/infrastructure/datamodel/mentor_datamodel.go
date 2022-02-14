package datamodel

import (
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
)

type Mentor struct {
	MentorID     string
	UserID       string
	Title        string
	MainImg      string
	SubImg       string
	Category     string
	Detail       string
	CreatedAt    time.Time
	Plans        []mentordm.Plan
	MentorSkills []mentordm.MentorSkill
}

type Plan struct {
	PlanID     string
	MentorID   string
	Title      string
	Category   string
	Tag        string
	Detail     string
	PlanType   uint16
	Price      uint16
	PlanStatus uint16
	CreatedAt  time.Time
}

type MentorSkill struct {
	MentorSkillID   string
	MentorID        string
	Tag             string
	Assessment      uint16
	ExperienceYears uint16
	CreatedAt       time.Time
}
