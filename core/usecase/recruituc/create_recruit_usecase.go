package recruituc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type CreateRecruitUsecase interface {
	Create(
		userID string,
		title string,
		budget uint32,
		recruitType uint16,
		detail string,
		recruitStatus uint16,
	) error
}

type CreateRecruitUsecaseImpl struct {
	recruitRepo recruitdm.RecruitRepository
}

func NewCreateRecruitUsecase(
	recruitRepo recruitdm.RecruitRepository,
) CreateRecruitUsecase {
	return &CreateRecruitUsecaseImpl{
		recruitRepo: recruitRepo,
	}
}

func (ru *CreateRecruitUsecaseImpl) Create(
	userID string,
	title string,
	budget uint32,
	recruitType uint16,
	detail string,
	recruitStatus uint16,
) error {
	userIDIns, err := userdm.NewUserIDByVal(userID)
	if err != nil {
		return err
	}
	recruitTypeIns, err := recruitdm.NewRecruitType(recruitType)
	if err != nil {
		return err
	}
	recruitStatusIns, err := recruitdm.NewRecruitStatus(recruitStatus)
	if err != nil {
		return err
	}
	recruitID := recruitdm.NewRecruitID()
	recruit, err := recruitdm.NewRecruit(
		recruitID,
		userIDIns,
		title,
		budget,
		recruitTypeIns,
		detail,
		recruitStatusIns,
	)
	if err != nil {
		return err
	}

	// メンター募集生成時は、「下書き」もしくは「公開」のみ可能
	if !recruit.IsCreate() {
		return xerrors.New("RecruitStatus must be Draft or Published when create")
	}

	return ru.recruitRepo.Create(recruit)
}
