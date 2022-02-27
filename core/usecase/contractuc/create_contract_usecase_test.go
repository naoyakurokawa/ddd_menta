package contractuc

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/naoyakurokawa/ddd_menta/core/domain/contractdm/mock_contractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	mockMentor "github.com/naoyakurokawa/ddd_menta/core/domain/mentordm/mock_mentordm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	asserts := assert.New(t)
	ctrl := gomock.NewController(t)
	mockContractRepository := mock.NewMockContractRepository(ctrl)
	mockContractRepository.EXPECT().Create(gomock.Any()).Return(nil)

	for _, td := range []struct {
		title  string
		input  uint16
		output string
	}{
		{
			title:  "PlanStatusがActiveのとき_エラーが発生しないこと",
			input:  uint16(1),
			output: "",
		},
		{
			title:  "PlanStatusがBusyのとき_エラーが発生すること",
			input:  uint16(2),
			output: "This plan is not active",
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

			planStatus, err := mentordm.NewPlanStatus(td.input)
			require.NoError(t, err)
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

			mockMentorRepository := mockMentor.NewMockMentorRepository(ctrl)
			mockMentorRepository.EXPECT().FindByID(gomock.Any()).Return(mentor, nil)

			contractUsecase := NewCreateContractUsecase(mockContractRepository, mockMentorRepository)
			err = contractUsecase.Create(
				up.userID.String(),
				mp.mentorID.String(),
				mp.planID.String(),
			)

			strErr := ""
			if err != nil {
				strErr = err.Error()
			}

			asserts.Equal(td.output, strErr)
		})
	}

}
