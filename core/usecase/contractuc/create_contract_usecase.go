package contractuc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/contractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type CreateContractUsecase interface {
	Create(
		userID string,
		mentorID string,
		planID string,
	) error
}

type CreateContractUsecaseImpl struct {
	contractRepo contractdm.ContractRepository
	mentorRepo   mentordm.MentorRepository
}

func NewCreateContractUsecase(
	contractRepo contractdm.ContractRepository,
	mentorRepo mentordm.MentorRepository,
) CreateContractUsecase {
	return &CreateContractUsecaseImpl{
		contractRepo: contractRepo,
		mentorRepo:   mentorRepo,
	}
}

func (cu *CreateContractUsecaseImpl) Create(
	userID string,
	mentorID string,
	planID string,
) error {
	userIDIns, err := userdm.NewUserIDByVal(userID)
	if err != nil {
		return err
	}
	mentorIDIns, err := mentordm.NewMentorIDByVal(mentorID)
	if err != nil {
		return err
	}
	planIDIns, err := mentordm.NewPlanIDByVal(planID)
	if err != nil {
		return err
	}
	mentorDomainService := mentordm.NewMentorDomainService(cu.mentorRepo)
	if !mentorDomainService.IsActivePlan(mentorIDIns, planIDIns) {
		return xerrors.New("This plan is not active")
	}
	// mentor := repoimpl.
	// CheckRequestPlanStatus
	//メンティーによる契約リクエスト時のStatusは未承認
	status := contractdm.Unapproved

	contract, err := contractdm.NewContract(
		userIDIns,
		mentorIDIns,
		planIDIns,
		status,
	)
	if err != nil {
		return err
	}

	err = cu.contractRepo.Create(contract)
	if err != nil {
		return err
	}

	return nil
}
