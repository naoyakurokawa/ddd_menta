package repoimpl

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/personalcontractdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/datamodel"
)

type PersonalContractRepositoryImpl struct {
	conn *gorm.DB
}

func NewPersonalContractRepositoryImpl(conn *gorm.DB) personalcontractdm.PersonalContractRepository {
	return &PersonalContractRepositoryImpl{conn: conn}
}

func (pr *PersonalContractRepositoryImpl) Create(personalContract *personalcontractdm.PersonalContract) error {
	var p datamodel.PersonalContract
	p.PersonalContractID = personalContract.PersonalContractID().String()
	p.SuggestionID = personalContract.SuggestionID().String()
	p.PersonalContractStatus = personalContract.PersonalContractStatus().Uint16()
	p.CreatedAt = time.Time(personalContract.CreatedAt())
	p.UpdatedAt = time.Time(personalContract.UpdatedAt())

	if err := pr.conn.Create(&p).Error; err != nil {
		return err
	}

	return nil
}

func (cr *PersonalContractRepositoryImpl) FindByID(
	personalContractID personalcontractdm.PersonalContractID,
) (*personalcontractdm.PersonalContract, error) {
	dataModelPersonalContract := datamodel.PersonalContract{}
	if err := cr.conn.Where("personal_contract_id = ?", string(personalContractID)).Find(&dataModelPersonalContract).Error; err != nil {
		return nil, err
	}

	return personalcontractdm.Reconstruct(
		dataModelPersonalContract.PersonalContractID,
		dataModelPersonalContract.SuggestionID,
		dataModelPersonalContract.PersonalContractStatus,
	)

}
