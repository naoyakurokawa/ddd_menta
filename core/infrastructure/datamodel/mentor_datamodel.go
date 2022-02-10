package datamodel

import (
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
)

type Mentor struct {
	UserID       string
	MentorID     string
	Title        string
	MainImg      string
	SubImg       string
	Category     string
	Detial       string
	CreatedAt    time.Time
	Plans        []mentordm.Plan
	MentorSkills []mentordm.MentorSkill
}

type Plan struct {
	PlanID     string
	Title      string
	Category   string
	Tag        string
	Detial     string
	PlanType   uint16
	Price      uint16
	PlanStatus uint16
	CreatedAt  time.Time
}

type MentorSkill struct {
	MentorSkillID   string
	Tag             string
	Assessment      uint16
	ExperienceYears uint16
	CreatedAt       time.Time
}
