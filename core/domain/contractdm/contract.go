package contractdm

import (
	"strconv"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type Contract struct {
	contractID     ContractID
	userID         userdm.UserID
	mentorID       mentordm.MentorID
	planID         mentordm.PlanID
	contractStatus ContractStatus
	createdAt      sharedvo.CreatedAt
	updatedAt      sharedvo.UpdatedAt
}

func NewContract(
	contractID ContractID,
	userID userdm.UserID,
	mentorID mentordm.MentorID,
	planID mentordm.PlanID,
	contractStatus ContractStatus,
) (*Contract, error) {

	contract := &Contract{
		contractID:     contractID,
		userID:         userID,
		mentorID:       mentorID,
		planID:         planID,
		contractStatus: contractStatus,
		createdAt:      sharedvo.NewCreatedAt(),
		updatedAt:      sharedvo.NewUpdatedAt(),
	}

	return contract, nil
}

func Reconstruct(
	contractID string,
	userID string,
	mentorID string,
	planID string,
	contractStatus uint16,
) (*Contract, error) {
	castedContractID, err := NewContractIDByVal(contractID)
	if err != nil {
		return nil, xerrors.New("error NewContractIDByVal")
	}
	castedUserID, err := userdm.NewUserIDByVal(userID)
	if err != nil {
		return nil, xerrors.New("error NewUserIDByVal")
	}
	castedMentorID, err := mentordm.NewMentorIDByVal(mentorID)
	if err != nil {
		return nil, xerrors.New("error NewMentorIDByVal")
	}
	castedPlanID, err := mentordm.NewPlanIDByVal(planID)
	if err != nil {
		return nil, xerrors.New("error NewMentorIDByVal")
	}
	ContractStatusIns, err := NewContractStatus(contractStatus)
	if err != nil {
		return nil, xerrors.New("error NewContractStatus")
	}

	contract := &Contract{
		contractID:     castedContractID,
		userID:         castedUserID,
		mentorID:       castedMentorID,
		planID:         castedPlanID,
		contractStatus: ContractStatusIns,
		createdAt:      sharedvo.NewCreatedAt(),
		updatedAt:      sharedvo.NewUpdatedAt(),
	}

	return contract, nil
}

func (c *Contract) ContractID() ContractID {
	return c.contractID
}

func (c *Contract) UserID() userdm.UserID {
	return c.userID
}

func (c *Contract) MentorID() mentordm.MentorID {
	return c.mentorID
}

func (c *Contract) PlanID() mentordm.PlanID {
	return c.planID
}

func (c *Contract) ContractStatus() ContractStatus {
	return c.contractStatus
}

func (c *Contract) CreatedAt() sharedvo.CreatedAt {
	return c.createdAt
}

func (c *Contract) UpdatedAt() sharedvo.UpdatedAt {
	return c.updatedAt
}

func StrCastUint(str string) (uint16, error) {
	ui, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(ui), nil
}

func (c *Contract) CanUpdateContractStatus(
	requestContractStatus ContractStatus,
) bool {
	return requestContractStatus.Uint16() == c.contractStatus.Uint16()+1
}
