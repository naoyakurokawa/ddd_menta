package contractdm

import (
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type contractParams struct {
	contractID ContractID
	userID     userdm.UserID
	mentorID   mentordm.MentorID
	planID     mentordm.PlanID
	status     Status
	createdAt  time.Time
	updatedAt  time.Time
}

var (
	cp contractParams
)

func setup() error {
	contractID := NewContractID()
	userID := userdm.NewUserID()
	mentorID := mentordm.NewMentorID()
	planID := mentordm.NewPlanID()
	status, err := NewStatus(uint16(1))
	if err != nil {
		return xerrors.New("error NewStatus")
	}
	cp = contractParams{
		contractID,
		userID,
		mentorID,
		planID,
		status,
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	return nil
}
