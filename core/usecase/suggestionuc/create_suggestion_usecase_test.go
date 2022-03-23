package suggestionuc

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	mentorMock "github.com/naoyakurokawa/ddd_menta/core/domain/mentordm/mock_mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	recruitMock "github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm/mock_recruitdm"
	suggestionMock "github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm/mock_suggestiondm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func TestCreate(t *testing.T) {
	asserts := assert.New(t)
	type fields struct {
		suggestionRepository *suggestionMock.MockSuggestionRepository
		mentorRepository     *mentorMock.MockMentorRepository
		recruitRepository    *recruitMock.MockRecruitRepository
	}

	for _, td := range []struct {
		title            string
		mentorID         string
		recruitID        string
		price            uint32
		suggestionType   uint16
		detail           string
		suggestionStatus uint16
		expectedErr      error
		prepareMock      func(f *fields) error
	}{
		{
			title:            "MentorSkillが5つ以上かつ_Recruitが公開状態の場合_エラーが発生しないこと",
			mentorID:         sp.mentorID.String(),
			recruitID:        sp.recruitID.String(),
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusUnapproved.Uint16(),
			expectedErr:      nil,
			prepareMock: func(f *fields) error {
				mentorSkill, err := mentordm.NewMentorSkill(
					mp.mentorSkillID,
					mp.mentorTag,
					mp.mentorAssessment,
					mp.mentorExperienceYears,
				)
				if err != nil {
					return err
				}
				var mentorSkills []mentordm.MentorSkill
				for i := 0; i < 5; i++ {
					mentorSkills = append(mentorSkills, *mentorSkill)
				}

				mentor, err := mentordm.NewMentor(
					mp.mentorID,
					mp.userID,
					mp.title,
					mp.mainImg,
					mp.subImg,
					mp.category,
					mp.detial,
					mentorSkills,
					mentorPlans,
				)
				if err != nil {
					return err
				}
				f.mentorRepository.EXPECT().FindByID(gomock.Any()).Return(mentor, nil)

				recruit, err := recruitdm.NewRecruit(
					rp.recruitID,
					rp.userID,
					rp.title,
					rp.budget,
					rp.recruitTypeOnce,
					rp.detail,
					rp.recruitStatusPublished,
				)
				if err != nil {
					return err
				}
				f.recruitRepository.EXPECT().FetchByID(gomock.Any()).Return(recruit, nil)

				f.suggestionRepository.EXPECT().Create(gomock.Any()).Return(nil)

				return nil
			},
		},
		{
			title:            "MentorSkillが5つ未満の場合_エラーが発生すること",
			mentorID:         sp.mentorID.String(),
			recruitID:        sp.recruitID.String(),
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusUnapproved.Uint16(),
			expectedErr:      xerrors.New("Can be suggested if you have 5 or more mentor skills"),
			prepareMock: func(f *fields) error {
				mentorSkill, err := mentordm.NewMentorSkill(
					mp.mentorSkillID,
					mp.mentorTag,
					mp.mentorAssessment,
					mp.mentorExperienceYears,
				)
				if err != nil {
					return err
				}
				var mentorSkills []mentordm.MentorSkill
				for i := 0; i < 4; i++ {
					mentorSkills = append(mentorSkills, *mentorSkill)
				}

				mentor, err := mentordm.NewMentor(
					mp.mentorID,
					mp.userID,
					mp.title,
					mp.mainImg,
					mp.subImg,
					mp.category,
					mp.detial,
					mentorSkills,
					mentorPlans,
				)
				if err != nil {
					return err
				}
				f.mentorRepository.EXPECT().FindByID(gomock.Any()).Return(mentor, nil)

				recruit, err := recruitdm.NewRecruit(
					rp.recruitID,
					rp.userID,
					rp.title,
					rp.budget,
					rp.recruitTypeOnce,
					rp.detail,
					rp.recruitStatusPublished,
				)
				if err != nil {
					return err
				}
				f.recruitRepository.EXPECT().FetchByID(gomock.Any()).Return(recruit, nil)

				return nil
			},
		},
		{
			title:            "Recruitが公開以外の状態の場合_エラーが発生すること",
			mentorID:         sp.mentorID.String(),
			recruitID:        sp.recruitID.String(),
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusUnapproved.Uint16(),
			expectedErr:      xerrors.New("This recruit is not active"),
			prepareMock: func(f *fields) error {
				recruit, err := recruitdm.NewRecruit(
					rp.recruitID,
					rp.userID,
					rp.title,
					rp.budget,
					rp.recruitTypeOnce,
					rp.detail,
					rp.recruitStatusDraft,
				)
				if err != nil {
					return err
				}
				f.recruitRepository.EXPECT().FetchByID(gomock.Any()).Return(recruit, nil)

				return nil
			},
		},
		{
			title:            "mentorIDが空の場合_エラーが発生すること",
			mentorID:         "",
			recruitID:        sp.recruitID.String(),
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusUnapproved.Uint16(),
			expectedErr:      xerrors.New("error NewMentorIDByVal"),
			prepareMock:      nil,
		},
		{
			title:            "recruitIDが空の場合_エラーが発生すること",
			mentorID:         sp.mentorID.String(),
			recruitID:        "",
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusUnapproved.Uint16(),
			expectedErr:      xerrors.New("error NewRecruitIDByVal"),
			prepareMock:      nil,
		},
		{
			title:            "priceが1000未満の場合_エラーが発生すること",
			mentorID:         sp.mentorID.String(),
			recruitID:        sp.recruitID.String(),
			price:            0,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusUnapproved.Uint16(),
			expectedErr:      xerrors.New("price more than ¥1000"),
			prepareMock:      nil,
		},
		{
			title:            "suggestionTypeが0の場合_エラーが発生すること",
			mentorID:         sp.mentorID.String(),
			recruitID:        sp.recruitID.String(),
			price:            sp.price,
			suggestionType:   0,
			detail:           sp.detail,
			suggestionStatus: sp.suggestionStatusUnapproved.Uint16(),
			expectedErr:      xerrors.New("SuggestionType must be 1 or 2"),
			prepareMock:      nil,
		},
		{
			title:            "detailが空の場合_エラーが発生すること",
			mentorID:         sp.mentorID.String(),
			recruitID:        sp.recruitID.String(),
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           "",
			suggestionStatus: sp.suggestionStatusUnapproved.Uint16(),
			expectedErr:      xerrors.New("detail must not be empty"),
			prepareMock:      nil,
		},
		{
			title:            "suggestionStatusが0の場合_エラーが発生すること",
			mentorID:         sp.mentorID.String(),
			recruitID:        sp.recruitID.String(),
			price:            sp.price,
			suggestionType:   sp.suggestionTypeOnce.Uint16(),
			detail:           sp.detail,
			suggestionStatus: 0,
			expectedErr:      xerrors.New("SuggestionStatus must be 1 or 2 or 3"),
			prepareMock:      nil,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				suggestionRepository: suggestionMock.NewMockSuggestionRepository(ctrl),
				mentorRepository:     mentorMock.NewMockMentorRepository(ctrl),
				recruitRepository:    recruitMock.NewMockRecruitRepository(ctrl),
			}
			if td.prepareMock != nil {
				if err := td.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			createSuggestionUsecase := NewCreateSuggestionUsecase(
				f.suggestionRepository,
				f.mentorRepository,
				f.recruitRepository,
			)

			err := createSuggestionUsecase.Create(
				td.mentorID,
				td.recruitID,
				td.price,
				td.suggestionType,
				td.detail,
				td.suggestionStatus,
			)

			if td.expectedErr != nil {
				require.Error(t, err)
				asserts.Equal(err.Error(), td.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}

}
