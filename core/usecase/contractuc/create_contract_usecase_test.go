package contractuc

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/naoyakurokawa/ddd_menta/core/domain/contractdm/mock_contractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	mockMentor "github.com/naoyakurokawa/ddd_menta/core/domain/mentordm/mock_mentordm"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type fields struct {
		contractRepository *mock.MockContractRepository
		mentorRepository   *mockMentor.MockMentorRepository
	}
	var mentorSkills []mentordm.MentorSkill

	for _, td := range []struct {
		title          string
		uintPlanStatus uint16
		userID         string
		mentorID       string
		planID         string
		expectedErr    error
		prepareMock    func(f *fields) error
	}{
		{
			title:       "PlanStatusがActiveのとき_エラーが発生しないこと",
			userID:      up.userID.String(),
			mentorID:    mp.mentorID.String(),
			planID:      mp.planID.String(),
			expectedErr: nil,
			prepareMock: func(f *fields) error {
				planStatus, err := mentordm.NewPlanStatus(1)
				if err != nil {
					return err
				}
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
				if err != nil {
					return err
				}
				var plans []mentordm.Plan
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
				if err != nil {
					return err
				}
				f.contractRepository.EXPECT().Create(gomock.Any()).Return(nil)
				f.mentorRepository.EXPECT().FindByID(gomock.Any()).Return(mentor, nil)

				return nil
			},
		},
		{
			title:       "UserIDが空の時_エラーが発生すること",
			userID:      "",
			mentorID:    mp.mentorID.String(),
			planID:      mp.planID.String(),
			expectedErr: errors.New("error NewUserIDByVal"),
			prepareMock: nil,
		},
		{
			title:       "MentorIDが空の時_エラーが発生すること",
			userID:      up.userID.String(),
			mentorID:    "",
			planID:      mp.planID.String(),
			expectedErr: errors.New("error NewMentorIDByVal"),
			prepareMock: nil,
		},
		{
			title:       "PlanIDが空の時_エラーが発生すること",
			userID:      up.userID.String(),
			mentorID:    mp.planID.String(),
			planID:      "",
			expectedErr: errors.New("error NewPlanIDByVal"),
			prepareMock: nil,
		},
		{
			title:       "PlanStatusがBusyのとき_エラーが発生すること",
			userID:      up.userID.String(),
			mentorID:    mp.mentorID.String(),
			planID:      mp.planID.String(),
			expectedErr: errors.New("This plan is not active"),
			prepareMock: func(f *fields) error {
				planStatus, err := mentordm.NewPlanStatus(2)
				if err != nil {
					return err
				}
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
				if err != nil {
					return err
				}
				var plans []mentordm.Plan
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
				if err != nil {
					return err
				}
				f.mentorRepository.EXPECT().FindByID(gomock.Any()).Return(mentor, nil)

				return nil
			},
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				contractRepository: mock.NewMockContractRepository(ctrl),
				mentorRepository:   mockMentor.NewMockMentorRepository(ctrl),
			}
			if td.prepareMock != nil {
				if err := td.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}
			contractUsecase := NewCreateContractUsecase(f.contractRepository, f.mentorRepository)
			err := contractUsecase.Create(
				td.userID,
				td.mentorID,
				td.planID,
			)

			if td.expectedErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
