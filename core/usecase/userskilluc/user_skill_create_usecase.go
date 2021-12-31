package userskilluc

import (
	"strconv"

	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userskilldm"
)

// UserSkillUsecase userskill usecaseのinterface
type UserSkillCreateUsecase interface {
	Create(userID string, tag []string, assessment []string, experienceYears []string) ([]*userskilldm.UserSkill, error)
}

type UserSkillCreateUsecaseImpl struct {
	userSkillRepo userskilldm.UserSkillRepository
}

// userSkill usecaseのコンストラクタ
func NewUserSkillCreateUsecase(userSkillRepo userskilldm.UserSkillRepository) UserSkillCreateUsecase {
	return &UserSkillCreateUsecaseImpl{userSkillRepo: userSkillRepo}
}

// Create userを保存するときのユースケース
func (uu *UserSkillCreateUsecaseImpl) Create(userID string, tag []string, assessment []string, experienceYears []string) ([]*userskilldm.UserSkill, error) {
	// userCareers := []userdm.UserCareer{}
	userSkills := make([]*userskilldm.UserSkill, len(tag))
	for i := 0; i < len(tag); i++ {
		userSkillID, err := userskilldm.NewUserSkillID()
		if err != nil {
			return nil, err
		}
		userID := userdm.UserIDType(userID)
		assessment, err := userskilldm.AssessmentCastUint(assessment[i])
		if err != nil {
			return nil, err
		}
		experienceYears, _ := strconv.Atoi(experienceYears[i])
		experienceYearsIns, err := userskilldm.NewExperienceYears(experienceYears)
		if err != nil {
			return nil, err
		}
		userSkill, err := userskilldm.NewUserSkill(userSkillID, userID, tag[i], assessment, experienceYearsIns)
		if err != nil {
			return nil, err
		}
		userSkills[i] = userSkill
	}

	//最終的にinfraのCreateメソッドを実行することになる
	_, err := uu.userSkillRepo.Create(userSkills)
	if err != nil {
		return nil, err
	}

	return userSkills, nil
}
