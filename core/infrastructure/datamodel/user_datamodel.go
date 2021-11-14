package datamodel

import (
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

type User struct {
	UserID      userdm.UserID
	Name        string
	Email       userdm.Email
	Password    string
	Profile     string
	CreatedAt   time.Time
	UserCareers []userdm.UserCareer
}

type UserCareer struct {
	UserCareerID userdm.UserCareerID
	UserID       userdm.UserID
	From         time.Time
	To           time.Time
	Detail       string
	CreatedAt    time.Time
}

type UserCareers []UserCareer
