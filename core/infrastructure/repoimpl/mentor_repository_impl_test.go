package repoimpl

import (
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	//ユーザー作成
	err := setupUser()
	require.NoError(t, err)
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
	err = userRepository.Create(user)
	require.NoError(t, err)

	// メンター作成
	err = setupMentor()
	require.NoError(t, err)
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

	selectedMentor, err := mentorRepository.FindByID(mp.mentorID)
	require.NoError(t, err)

	// 検証 (作成したメンターがDB登録されているか)
	expected := mp.mentorID
	actual := selectedMentor.MentorID()
	assert.Equal(t, expected, actual)
}
