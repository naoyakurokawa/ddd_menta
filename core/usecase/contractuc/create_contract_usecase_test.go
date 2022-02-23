package contractuc

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/naoyakurokawa/ddd_menta/core/domain/contractdm/mock_contractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	mockMentor "github.com/naoyakurokawa/ddd_menta/core/domain/mentordm/mock_mentordm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func TestCreate(t *testing.T) {
	asserts := assert.New(t)
	for _, td := range []struct {
		title string
		input uint16
		err   error
	}{
		{
			title: "PlanStatusがActiveのとき_エラーが発生しないこと",
			input: uint16(1),
			err:   nil,
		},
		{
			title: "PlanStatusがBusyのとき_エラーが発生すること",
			input: uint16(2),
			err:   xerrors.New("This plan is not active"),
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			var (
				mentorSkills []mentordm.MentorSkill
				plans        []mentordm.Plan
			)
			setupUser()
			setupMentor()
			setupContract()

			ctrl := gomock.NewController(t)
			mockContractRepository := mock.NewMockContractRepository(ctrl)
			mockContractRepository.EXPECT().Create(gomock.Any()).Return(nil)
			mockMentorRepository := mockMentor.NewMockMentorRepository(ctrl)
			mentorSkill, err := mentordm.NewMentorSkill(
				mp.mentorSkillID,
				mp.mentorTag,
				mp.mentorAssessment,
				mp.mentorExperienceYears,
			)
			require.NoError(t, err)
			mentorSkills = append(mentorSkills, *mentorSkill)
			planStatus, err := mentordm.NewPlanStatus(td.input)
			plan, err := mentordm.NewPlan(
				mp.planID,
				mp.planTitle,
				mp.category,
				mp.planTag,
				mp.planDetial,
				mp.planType,
				mp.planPrice,
				planStatus,
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
			mockMentorRepository.EXPECT().FindByID(gomock.Any()).Return(mentor, nil)
			contractUsecase := NewCreateContractUsecase(mockContractRepository, mockMentorRepository)
			err = contractUsecase.Create(
				string(up.userID),
				string(mp.mentorID),
				string(mp.planID),
			)
			asserts.Equal(err, td.err)
		})
	}

}
