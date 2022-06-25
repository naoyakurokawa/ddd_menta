package suggestiondm

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
)

type suggestionParams struct {
	suggestionID               SuggestionID
	mentorID                   mentordm.MentorID
	recruitID                  recruitdm.RecruitID
	price                      uint32
	suggestionTypeOnce         SuggestionType
	suggestionTypeSubscription SuggestionType
	detail                     string
	suggestionStatusUnapproved SuggestionStatus
	suggestionStatusApproval   SuggestionStatus
	suggestionStatusTerminated SuggestionStatus
	createdAt                  time.Time
	updatedAt                  time.Time
}

var (
	sp suggestionParams
)

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		fmt.Printf("%+v", err)
		return
	}
	os.Exit(m.Run())
}

func setup() error {
	suggestionID := NewSuggestionID()
	mentorID := mentordm.NewMentorID()
	recruitID := recruitdm.NewRecruitID()
	sp = suggestionParams{
		suggestionID,
		mentorID,
		recruitID,
		5000,
		Once,
		Subscription,
		"DDDを設計から開発までサポートします",
		Unapproved,
		Approval,
		Terminated,
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	return nil
}
