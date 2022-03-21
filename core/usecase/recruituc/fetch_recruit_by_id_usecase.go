package recruituc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
)

type FetchRecruitByIdUsecase interface {
	FetchByID(recruitID recruitdm.RecruitID) (*recruitdm.Recruit, error)
}

type FetchRecruitByIdUsecaseImpl struct {
	recruitRepo recruitdm.RecruitRepository
}

func NewMFetchRecruitByIdUsecase(recruitRepo recruitdm.RecruitRepository) FetchRecruitByIdUsecase {
	return &FetchRecruitByIdUsecaseImpl{recruitRepo: recruitRepo}
}

func (ru *FetchRecruitByIdUsecaseImpl) FetchByID(recruitID recruitdm.RecruitID) (*recruitdm.Recruit, error) {
	return ru.recruitRepo.FetchByID(recruitID)
}
