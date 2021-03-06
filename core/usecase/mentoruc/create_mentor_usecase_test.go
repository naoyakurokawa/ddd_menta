package mentoruc

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	mock "github.com/naoyakurokawa/ddd_menta/core/domain/mentordm/mock_mentordm"
)

func TestCreate(t *testing.T) {
	err := setupUser()
	require.NoError(t, err)
	err = setupMentor()
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	mockMentorRepository := mock.NewMockMentorRepository(ctrl)
	mockMentorRepository.EXPECT().Create(gomock.Any()).Return(nil)
	mentorUsecase := NewCreateMentorUsecase(mockMentorRepository)
	err = mentorUsecase.Create(
		mp.userID.String(),
		mp.title,
		mp.mainImg,
		mp.subImg,
		mp.category,
		mp.detail,
		mp.mentorSkills,
		mp.plans,
	)

	// 検証 エラーが発生しないこと
	require.NoError(t, err)
}
