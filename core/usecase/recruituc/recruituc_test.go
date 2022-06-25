package recruituc

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

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

var (
	rp recruitParams
)

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		fmt.Printf("%+v", err)
		return
	}
	os.Exit(m.Run())
}

func setup() error {
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
