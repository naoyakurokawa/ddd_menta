package mentoruc

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	mock "github.com/naoyakurokawa/ddd_menta/core/domain/mentordm/mock_mentordm"
)

func TestCreate(t *testing.T) {
	setupUser()
	setupMentor()

	ctrl := gomock.NewController(t)
	mockMentorRepository := mock.NewMockMentorRepository(ctrl)
	mockMentorRepository.EXPECT().Create(gomock.Any()).Return(nil)
	mentorUsecase := NewCreateMentorUsecase(mockMentorRepository)
	err := mentorUsecase.Create(
		string(mp.userID),
		mp.title,
		mp.mainImg,
		mp.subImg,
		mp.category,
		mp.detail,
		mp.mentorTag,
		mp.mentorAssessment,
		mp.mentorExperienceYears,
		mp.plans,
	)

	// 検証 エラーが発生しないこと
	require.NoError(t, err)
}
