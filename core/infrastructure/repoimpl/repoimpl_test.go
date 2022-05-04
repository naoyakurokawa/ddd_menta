package repoimpl

import (
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/contractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/personalcontractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type userParams struct {
	userID    userdm.UserID
	name      string
	email     sharedvo.Email
	password  sharedvo.Password
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

type contractParams struct {
	contractID     contractdm.ContractID
	userID         userdm.UserID
	mentorID       mentordm.MentorID
	planID         mentordm.PlanID
	contractStatus contractdm.ContractStatus
	createdAt      time.Time
	updatedAt      time.Time
}

type recruitParams struct {
	recruitID     recruitdm.RecruitID
	userID        userdm.UserID
	title         string
	budget        uint32
	recruitType   recruitdm.RecruitType
	detail        string
	recruitStatus recruitdm.RecruitStatus
	createdAt     time.Time
	updatedAt     time.Time
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

type personalContractParams struct {
	personalContractID       personalcontractdm.PersonalContractID
	suggestionID             suggestiondm.SuggestionID
	unapprovedStatus         personalcontractdm.PersonalContractStatus
	underContractStatus      personalcontractdm.PersonalContractStatus
	terminatedContractStatus personalcontractdm.PersonalContractStatus
	createdAt                time.Time
	updatedAt                time.Time
}

var (
	up           userParams
	mp           mentorParams
	cp           contractParams
	rp           recruitParams
	sp           suggestionParams
	pp           personalContractParams
	userCareers  []userdm.UserCareer
	userSkills   []userdm.UserSkill
	mentorSkills []mentordm.MentorSkill
	mentorPlans  []mentordm.Plan
)

func setupUser() error {
	//ユーザー
	userID := userdm.NewUserID()
	email, err := sharedvo.NewEmail("test@gmail.com")
	if err != nil {
		return xerrors.New("error NewEmail")
	}
	password, err := sharedvo.NewPassword("test12345678")
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

func setupContract() error {
	contractID := contractdm.NewContractID()
	userID := userdm.NewUserID()
	mentorID := mentordm.NewMentorID()
	planID := mentordm.NewPlanID()
	contractStatus, err := contractdm.NewContractStatus(uint16(1))
	if err != nil {
		return xerrors.New("error NewStatus")
	}
	cp = contractParams{
		contractID,
		userID,
		mentorID,
		planID,
		contractStatus,
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	return nil
}

func setupRecruit() error {
	recruitID := recruitdm.NewRecruitID()
	rp = recruitParams{
		recruitID,
		up.userID,
		"DDDの基礎を教えて下さい",
		5000,
		recruitdm.Once,
		"DDDによる開発をサポートしてもらいたく募集しました",
		recruitdm.Draft,
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

func setupPersonalContact() error {
	personalContactID := personalcontractdm.NewPersonalContractID()
	pp = personalContractParams{
		personalContactID,
		sp.suggestionID,
		personalcontractdm.Unapproved,
		personalcontractdm.UnderContract,
		personalcontractdm.TerminatedContract,
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	return nil
}
