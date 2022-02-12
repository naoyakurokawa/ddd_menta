package repoimpl

import (
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	//ユーザー作成
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
	if err != nil {
		t.Errorf("failed to NewUser: %v", err)
		return
	}
	userRepository := NewUserRepositoryImpl(db.NewDB())
	_, err = userRepository.Create(user)
	if err != nil {
		t.Errorf("failed to userRepository.Create: %v", err)
	}
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

	if err != nil {
		t.Errorf("failed to NewUser: %v", err)
		return
	}

	mentorRepository := NewMentorRepositoryImpl(db.NewDB())
	createdMentor, err := mentorRepository.Create(mentor)
	if err != nil {
		t.Errorf("failed to mentorRepository.Create: %v", err)
	}
	selectedMentor, err := mentorRepository.FindByID(createdMentor.MentorID())
	if err != nil {
		t.Errorf("failed to FindByID: %v", err)
	}
	if !createdMentor.MentorID().Equals(selectedMentor.MentorID()) {
		t.Errorf("failed to CreateUser")
	}
}
