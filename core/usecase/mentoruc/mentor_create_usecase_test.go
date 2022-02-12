package mentoruc

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	mock "github.com/naoyakurokawa/ddd_menta/core/domain/mentordm/mock_mentordm"
)

func TestCreate(t *testing.T) {
	//ユーザー作成
	setupUser()

	// メンター作成
	setupMentor()
	mentor, err := mentordm.NewMentor(
		mp.mentorID,
		mp.userID,
		mp.title,
		mp.mainImg,
		mp.subImg,
		mp.category,
		mp.detail,
		mentorSkills,
		plans,
	)
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	mockMentorRepository := mock.NewMockMentorRepository(ctrl)
	mockMentorRepository.EXPECT().Create(gomock.Any()).Return(mentor, nil)
	mentorUsecase := NewMentorCreateUsecase(mockMentorRepository)
	_, err = mentorUsecase.Create(
		string(mp.userID),
		mp.title,
		mp.mainImg,
		mp.subImg,
		mp.category,
		mp.detail,
		mp.mentorTag,
		mp.mentorAssessment,
		mp.mentorExperienceYears,
		mp.planTitle,
		mp.planCategory,
		mp.planTag,
		mp.planDetial,
		mp.planType,
		mp.planPrice,
		mp.planStatus,
	)

	// 検証 エラーが発生しないこと
	require.NoError(t, err)
}
