package repoimpl

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userskilldm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/datamodel"
)

type UserSkillRepositoryImpl struct {
	Conn *gorm.DB
}

func NewUserSkillRepositoryImpl(conn *gorm.DB) userskilldm.UserSkillRepository {
	return &UserSkillRepositoryImpl{Conn: conn}
}

func (ur *UserSkillRepositoryImpl) Create(userSkills []*userskilldm.UserSkill) ([]*userskilldm.UserSkill, error) {
	tx := ur.Conn.Begin()
	for i := 0; i < len(userSkills); i++ {
		userSkill := &datamodel.UserSkill{
			UserSkillID:     userskilldm.UserSkillID.Value(userSkills[i].UserSkillID()),
			UserID:          userdm.UserID.Value(userSkills[i].UserID()),
			Tag:             userSkills[i].Tag(),
			Assessment:      userSkills[i].Assessment(),
			ExperienceYears: uint16(userSkills[i].ExperienceYears()),
			CreatedAt:       userSkills[i].CreatedAt(),
		}
		if err := tx.Create(&userSkill).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return userSkills, nil
}
