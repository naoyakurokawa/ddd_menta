package repoimpl

import (
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type userParams struct {
	userID    userdm.UserID
	name      string
	email     userdm.Email
	password  userdm.Password
	profile   string
	createdAt time.Time
}

type mentorParams struct {
	userID                userdm.UserID
	mentorID              mentordm.MentorID
	title                 string
	mainImg               string
	subImg                string
	category              string
	detial                string
	mentorSkillID         mentordm.MentorSkillID
	mentorTag             string
	mentorAssessment      uint16
	mentorExperienceYears mentordm.ExperienceYears
	planID                mentordm.PlanID
	planTitle             string
	planCategory          string
	planTag               string
	planDetial            string
	planType              mentordm.PlanType
	planPrice             uint16
	planStatus            mentordm.PlanStatus
	createdAt             time.Time
}

var (
	up          userParams
	mp          mentorParams
	userCareers []userdm.UserCareer
	userSkills  []userdm.UserSkill
)

func setupUser() error {
	//ユーザー
	userID, err := userdm.NewUserID()
	if err != nil {
		return xerrors.New("error NewUserID")
	}
	email, err := userdm.NewEmail("test@gmail.com")
	if err != nil {
		return xerrors.New("error NewEmail")
	}
	password, err := userdm.NewPassword("test12345678")
	if err != nil {
		return xerrors.New("error NewPassword")
	}
	up = userParams{
		userID,
		"テストユーザー",
		email,
		password,
		"テストユーザーです",
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	return nil
}

func setupMentor() error {
	mentorID := mentordm.NewMentorID()
	mentorSkillID := mentordm.NewMentorSkillID()
	planID := mentordm.NewPlanID()
	experienceYears, err := mentordm.NewExperienceYears(uint16(5))
	if err != nil {
		return xerrors.New("error NewExperienceYears")
	}
	planType, err := mentordm.NewPlanType(uint16(2))
	if err != nil {
		return xerrors.New("error NewPlanType")
	}
	planStatus, err := mentordm.NewPlanStatus(uint16(1))
	if err != nil {
		return xerrors.New("error NewPlanStatus")
	}

	mp = mentorParams{
		up.userID,
		mentorID,
		"プログラミング全般のメンタリング",
		"/main.jpg",
		"/sub.jpg",
		"プログライミング",
		"設計・開発・テストの一覧をサポートできます",
		mentorSkillID,
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
