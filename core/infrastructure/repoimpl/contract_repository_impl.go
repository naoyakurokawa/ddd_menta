package repoimpl

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/contractdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/datamodel"
)

type ContractRepositoryImpl struct {
	conn *gorm.DB
}

func NewContractRepositoryImpl(conn *gorm.DB) contractdm.ContractRepository {
	return &ContractRepositoryImpl{conn: conn}
}

func (cr *ContractRepositoryImpl) Create(contract *contractdm.Contract) error {
	var c datamodel.Contract
	c.ContractID = string(contract.ContractID())
	c.UserID = string(contract.UserID())
	c.MentorID = string(contract.MentorID())
	c.PlanID = string(contract.PlanID())
	c.Status = uint16(contract.Status())
	c.CreatedAt = time.Time(contract.CreatedAt())
	c.UpdatedAt = time.Time(contract.UpdatedAt())

	if err := cr.conn.Create(&c).Error; err != nil {
		return err
	}

	return nil
}

func (cr *ContractRepositoryImpl) FindByID(contractID contractdm.ContractID) (*contractdm.Contract, error) {
	dataModeContract := &datamodel.Contract{}
	if err := cr.conn.Where("contract_id = ?", string(contractID)).Find(&dataModeContract).Error; err != nil {
		return nil, err
	}

	contract, err := contractdm.Reconstruct(
		dataModeContract.ContractID,
		dataModeContract.UserID,
		dataModeContract.MentorID,
		dataModeContract.PlanID,
		dataModeContract.Status,
	)
	if err != nil {
		return nil, err
	}

	return contract, nil
}
