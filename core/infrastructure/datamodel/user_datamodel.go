package datamodel

import (
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

type User struct {
	UserID      string
	Name        string
	Email       string
	Password    string
	Profile     string
	CreatedAt   time.Time
	UserCareers []userdm.UserCareer
	UserSkills  []userdm.UserSkill
}

type UserCareer struct {
	UserCareerID string
	UserID       string
	From         time.Time
	To           time.Time
	Detail       string
	CreatedAt    time.Time
}

type UserSkill struct {
	UserSkillID     string
	UserID          string
	Tag             string
	Assessment      uint16
	ExperienceYears uint16
	CreatedAt       time.Time
}
