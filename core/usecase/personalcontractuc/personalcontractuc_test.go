package personalcontractuc

import (
	"fmt"
	"testing"
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/personalcontractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
)

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
	sp suggestionParams
	pp personalContractParams
)

func TestMain(m *testing.M) {
	err := setupSuggestion()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	err = setupPersonalContact()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	m.Run()
}

func setupSuggestion() error {
	suggestionID := suggestiondm.NewSuggestionID()
	mentorID := mentordm.NewMentorID()
	recruitID := recruitdm.NewRecruitID()
	sp = suggestionParams{
		suggestionID,
		mentorID,
		recruitID,
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
