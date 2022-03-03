package repoimpl

import (
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/contractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContractRepoCreate(t *testing.T) {
	//ユーザー作成 (メンターとなるユーザー)
	setupUser()
	var (
		mentorSkills []mentordm.MentorSkill
		plans        []mentordm.Plan
	)
	user, err := userdm.NewUser(
		up.userID,
		up.name,
		up.email,
		up.password,
		up.profile,
		userCareers,
		userSkills,
	)
	require.NoError(t, err)
	userRepository := NewUserRepositoryImpl(db.NewDB())
	_, err = userRepository.Create(user)
	require.NoError(t, err)

	// メンター作成
	setupMentor()
	mentorSkill, err := mentordm.NewMentorSkill(
		mp.mentorSkillID,
		mp.mentorTag,
		mp.mentorAssessment,
		mp.mentorExperienceYears,
	)
	require.NoError(t, err)
	mentorSkills = append(mentorSkills, *mentorSkill)
	plan, err := mentordm.NewPlan(
		mp.planID,
		mp.planTitle,
		mp.category,
		mp.planTag,
		mp.planDetial,
		mp.planType,
		mp.planPrice,
		mp.planStatus,
	)
	require.NoError(t, err)
	plans = append(plans, *plan)
	mentor, err := mentordm.NewMentor(
		mp.mentorID,
		mp.userID,
		mp.title,
		mp.mainImg,
		mp.subImg,
		mp.category,
		mp.detial,
		mentorSkills,
		plans,
	)
	require.NoError(t, err)

	mentorRepository := NewMentorRepositoryImpl(db.NewDB())
	err = mentorRepository.Create(mentor)
	require.NoError(t, err)

	//ユーザー作成 (メンティーとなるユーザー)
	setupUser()
	user2, err := userdm.NewUser(
		up.userID,
		up.name,
		up.email,
		up.password,
		up.profile,
		userCareers,
		userSkills,
	)
	require.NoError(t, err)
	userRepository2 := NewUserRepositoryImpl(db.NewDB())
	_, err = userRepository2.Create(user2)
	require.NoError(t, err)

	//契約
	setupContract()
	contract, err := contractdm.NewContract(
		cp.contractID,
		up.userID,
		mp.mentorID,
		mp.planID,
		contractdm.Unapproved,
	)

	contractRepository := NewContractRepositoryImpl(db.NewDB())
	err = contractRepository.Create(contract)
	require.NoError(t, err)

	selectedContract, err := contractRepository.FindByID(cp.contractID)
	require.NoError(t, err)

	// 検証 (作成した契約がDB登録されているか)
	expected := cp.contractID
	actual := selectedContract.ContractID()
	assert.Equal(t, expected, actual)
}
