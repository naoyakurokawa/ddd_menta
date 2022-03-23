package suggestionuc

import (
	"fmt"
	"testing"
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

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
	planActiveStatus      mentordm.PlanStatus
	planBusyStatus        mentordm.PlanStatus
	createdAt             time.Time
}

type recruitParams struct {
	recruitID               recruitdm.RecruitID
	userID                  userdm.UserID
	title                   string
	budget                  uint32
	recruitTypeOnce         recruitdm.RecruitType
	recruitTypeSubscription recruitdm.RecruitType
	detail                  string
	recruitStatusDraft      recruitdm.RecruitStatus
	recruitStatusPublished  recruitdm.RecruitStatus
	recruitStatusTerminated recruitdm.RecruitStatus
	createdAt               time.Time
	updatedAt               time.Time
}

type suggestionParams struct {
	suggestionID               suggestiondm.SuggestionID
	mentorID                   mentordm.MentorID
	recruitID                  recruitdm.RecruitID
	price                      uint32
	suggestionTypeOnce         suggestiondm.SuggestionType
	suggestionTypeSubscription suggestiondm.SuggestionType
	detail                     string
	suggestionStatusUnapproved suggestiondm.SuggestionStatus
	suggestionStatusApproval   suggestiondm.SuggestionStatus
	suggestionStatusTerminated suggestiondm.SuggestionStatus
	createdAt                  time.Time
	updatedAt                  time.Time
}

var (
	mp          mentorParams
	rp          recruitParams
	sp          suggestionParams
	userCareers []userdm.UserCareer
	userSkills  []userdm.UserSkill
	mentorPlans []mentordm.Plan
)

func TestMain(m *testing.M) {
	err := setupMentor()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	err = setupRecruit()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	err = setupSuggestion()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	m.Run()
}

func setupMentor() error {
	userID := userdm.NewUserID()
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
	activePlanStatus, err := mentordm.NewPlanStatus(uint16(1))
	if err != nil {
		return xerrors.New("error NewPlanStatus")
	}
	busyPlanStatus, err := mentordm.NewPlanStatus(uint16(2))
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
		activePlanStatus,
		busyPlanStatus,
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	return nil
}

func setupRecruit() error {
	recruitID := recruitdm.NewRecruitID()
	userID := userdm.NewUserID()
	rp = recruitParams{
		recruitID,
		userID,
		"DDDの基礎を教えて下さい",
		5000,
		recruitdm.Once,
		recruitdm.Subscription,
		"DDDによる開発をサポートしてもらいたく募集しました",
		recruitdm.Draft,
		recruitdm.Published,
		recruitdm.Terminated,
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	return nil
}

func setupSuggestion() error {
	suggestionID := suggestiondm.NewSuggestionID()
	sp = suggestionParams{
		suggestionID,
		mp.mentorID,
		rp.recruitID,
		5000,
		suggestiondm.Once,
		suggestiondm.Subscription,
		"DDDの設計から開発までサポートします",
		suggestiondm.Unapproved,
		suggestiondm.Approval,
		suggestiondm.Terminated,
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	return nil
}
