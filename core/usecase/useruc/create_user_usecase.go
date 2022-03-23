package useruc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

// UserUsecase user usecaseのinterface
type CreateUserUsecase interface {
	Create(
		name string,
		email string,
		password string,
		profile string,
		userCareers []UserCareer,
		userSkills []UserSkill,
	) error
}

type UserCareer struct {
	From   string
	To     string
	Detail string
}

type UserSkill struct {
	Tag             string
	Assessment      string
	ExperienceYears string
}

type CreateUserUsecaseImpl struct {
	userRepo userdm.UserRepository
}

// user usecaseのコンストラクタ
func NewUserCreateUsecase(userRepo userdm.UserRepository) CreateUserUsecase {
	return &CreateUserUsecaseImpl{userRepo: userRepo}
}

// Create userを保存するときのユースケース
func (uu *CreateUserUsecaseImpl) Create(
	name string,
	email string,
	password string,
	profile string,
	userCareers []UserCareer,
	userSkills []UserSkill,
) error {
	var (
		initUserCareers []userdm.UserCareer
		initUserSkills  []userdm.UserSkill
	)

	userID := userdm.NewUserID()
	emailIns, err := userdm.NewEmail(email)
	if err != nil {
		return err
	}
	passwordIns, err := userdm.NewPassword(password)
	if err != nil {
		return err
	}

	user, err := userdm.NewUser(
		userID,
		name,
		emailIns,
		passwordIns,
		profile,
		initUserCareers,
		initUserSkills,
	)
	if err != nil {
		return err
	}

	for _, uc := range userCareers {
		user.AddUserCareer(
			uc.From,
			uc.To,
			uc.Detail,
		)
	}

	for _, us := range userSkills {
		uintUserAssessment, err := userdm.StrCastUint(us.Assessment)
		if err != nil {
			return err
		}

		uintUserExperienceYears, err := userdm.StrCastUint(us.ExperienceYears)
		if err != nil {
			return err
		}

		userExperienceYears, err := userdm.NewExperienceYears(uintUserExperienceYears)
		if err != nil {
			return err
		}

		user.AddUserSkill(
			us.Tag,
			uintUserAssessment,
			userExperienceYears,
		)
	}

	err = uu.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}
