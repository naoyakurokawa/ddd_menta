package contractdm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

func GenWhenCreate(
	contractID ContractID,
	userID userdm.UserID,
	mentorID mentordm.MentorID,
	planID mentordm.PlanID,
) (*Contract, error) {
	return NewContract(
		contractID,
		userID,
		mentorID,
		planID,
		Unapproved,
	)
}
