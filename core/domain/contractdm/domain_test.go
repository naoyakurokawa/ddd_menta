package contractdm

import (
	"testing"
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type contractParams struct {
	contractID     ContractID
	userID         userdm.UserID
	mentorID       mentordm.MentorID
	planID         mentordm.PlanID
	contractStatus ContractStatus
	createdAt      time.Time
	updatedAt      time.Time
}

var (
	cp contractParams
)

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() error {
	contractID := NewContractID()
	userID := userdm.NewUserID()
	mentorID := mentordm.NewMentorID()
	planID := mentordm.NewPlanID()
	contractStatus, err := NewContractStatus(uint16(1))
	if err != nil {
		return xerrors.New("error NewContractStatus")
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
