package repoimpl

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/datamodel"
	"golang.org/x/xerrors"
)

type RecruitRepositoryImpl struct {
	conn *gorm.DB
}

func NewRecruitRepositoryImpl(conn *gorm.DB) recruitdm.RecruitRepository {
	return &RecruitRepositoryImpl{conn: conn}
}

func (rr *RecruitRepositoryImpl) Create(recruit *recruitdm.Recruit) error {
	var r datamodel.Recruit
	r.RecruitID = recruit.RecruitID().String()
	r.UserID = recruit.UserID().String()
	r.Title = recruit.Title()
	r.Budget = recruit.Budget()
	r.RecruitType = recruit.RecruitType().Uint16()
	r.Detail = recruit.Detail()
	r.RecruitStatus = recruit.RecruitStatus().Uint16()
	if err := rr.conn.Create(&r).Error; err != nil {
		return xerrors.Errorf("fail to create recruit: %w", err)
	}

	return nil
}

func (rr *RecruitRepositoryImpl) FetchByID(recruitID recruitdm.RecruitID) (*recruitdm.Recruit, error) {
	dataModelRecruit := datamodel.Recruit{}
	if err := rr.conn.Where("recruit_id = ?", recruitID.String()).Find(&dataModelRecruit).Error; err != nil {
		return nil, xerrors.Errorf("fail to find by recruitID : %w", err)
	}

	return recruitdm.Reconstruct(
		dataModelRecruit.RecruitID,
		dataModelRecruit.UserID,
		dataModelRecruit.Title,
		dataModelRecruit.Budget,
		dataModelRecruit.RecruitType,
		dataModelRecruit.Detail,
		dataModelRecruit.RecruitStatus,
	)

}
