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
}

type UserCareer struct {
	UserCareerID string
	UserID       string
	From         time.Time
	To           time.Time
	Detail       string
	CreatedAt    time.Time
}
