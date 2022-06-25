package recruitdm

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

type recruitParams struct {
	recruitID               RecruitID
	userID                  userdm.UserID
	title                   string
	budget                  uint32
	recruitTypeOnce         RecruitType
	recruitTypeSubscription RecruitType
	detail                  string
	recruitStatusDraft      RecruitStatus
	recruitStatusPublished  RecruitStatus
	recruitStatusTerminated RecruitStatus
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
	recruitID := NewRecruitID()
	userID := userdm.NewUserID()
	rp = recruitParams{
		recruitID,
		userID,
		"DDDの基礎を教えて下さい",
		5000,
		Once,
		Subscription,
		"DDDによる開発をサポートしてもらいたく募集しました",
		Draft,
		Published,
		Terminated,
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	return nil
}
