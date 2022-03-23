package repoimpl

import (
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecruitRepoCreate(t *testing.T) {
	setupUser()
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

	setupRecruit()
	recruit, err := recruitdm.NewRecruit(
		rp.recruitID,
		rp.userID,
		rp.title,
		rp.budget,
		rp.recruitType,
		rp.detail,
		rp.recruitStatus,
	)
	require.NoError(t, err)
	recruitRepository := NewRecruitRepositoryImpl(db.NewDB())
	err = recruitRepository.Create(recruit)
	require.NoError(t, err)

	selectedRecruit, err := recruitRepository.FetchByID(rp.recruitID)
	require.NoError(t, err)

	// 検証 (作成したメンター募集がDB登録されているか)
	expected := rp.recruitID
	actual := selectedRecruit.RecruitID()
	assert.Equal(t, expected, actual)
}
