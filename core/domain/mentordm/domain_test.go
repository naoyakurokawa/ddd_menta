package mentordm

import (
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type mentorParams struct {
	userID                userdm.UserID
	mentorID              MentorID
	title                 string
	mainImg               string
	subImg                string
	category              string
	detial                string
	mentorSkillID         string
	mentorTag             string
	mentorAssessment      uint16
	mentorExperienceYears ExperienceYears
	planID                PlanID
	planTitle             string
	planCategory          string
	planTag               string
	planDetial            string
	planType              PlanType
	planPrice             uint16
	planStatus            PlanStatus
	createdAt             time.Time
}

var (
	mp mentorParams
)

func setup() error {
	userID, err := userdm.NewUserID()
	if err != nil {
		return xerrors.New("error NewUserID")
	}
	mentorID := NewMentorID()
	planID := NewPlanID()
	experienceYears, err := NewExperienceYears(uint16(5))
	if err != nil {
		return xerrors.New("error NewExperienceYears")
	}
	planType, err := NewPlanType(uint16(2))
	if err != nil {
		return xerrors.New("error NewPlanType")
	}
	planStatus, err := NewPlanStatus(uint16(1))
	if err != nil {
		return xerrors.New("error NewPlanStatus")
	}
	mp = mentorParams{
		userID,
		mentorID,
		"プログラミング全般のメンタリング",
		"/main.jpg",
		"/sub.jpg",
		"プログライミング",
		"設計・開発・テストの一覧をサポートできます",
		"mentorSkillID",
		"Golang",
		uint16(5),
		experienceYears,
		planID,
		"DDDのメンタリング",
		"設計",
		"DDD",
		"DDDの設計手法を学べます",
		planType,
		uint16(1000),
		planStatus,
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	return nil
}
