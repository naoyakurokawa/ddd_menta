package contractdm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

type Contract struct {
	contractID ContractID
	userID     userdm.UserID
	planID     mentordm.PlanID
	status     Status
	createdAt  sharedvo.CreatedAt
}

const titleMaxLength = 255
const detialMaxLength = 2000

func NewContract(
	userID userdm.UserID,
	planID mentordm.PlanID,
	status Status,
) (*Contract, error) {

	contractID := NewContractID()

	contract := &Contract{
		contractID: contractID,
		userID:     userID,
		planID:     planID,
		status:     status,
		createdAt:  sharedvo.NewCreatedAt(),
	}

	return contract, nil
}

func (c *Contract) ContractID() ContractID {
	return c.contractID
}

func (c *Contract) UserID() userdm.UserID {
	return c.userID
}

func (c *Contract) PlanID() mentordm.PlanID {
	return c.planID
}

func (c *Contract) Status() Status {
	return c.status
}
