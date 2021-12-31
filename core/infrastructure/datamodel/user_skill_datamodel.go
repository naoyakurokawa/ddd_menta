package datamodel

import (
	"time"
)

type UserSkill struct {
	UserSkillID     string
	UserID          string
	Tag             string
	Assessment      uint16
	ExperienceYears uint16
	CreatedAt       time.Time
}
