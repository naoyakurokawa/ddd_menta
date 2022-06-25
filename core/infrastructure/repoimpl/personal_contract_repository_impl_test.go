package repoimpl

import (
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/personalcontractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

type personalContractFields struct {
	personalContractID     string
	suggestionID           string
	personalContractStatus uint16
}

func TestPersonalContractRepoCreate(t *testing.T) {
	tx := db.NewDB().Begin()
	asserts := assert.New(t)
	for _, td := range []struct {
		title                   string
		preparePersonalContract func(s *personalContractFields) error
		expectedErr             error
		isLast                  bool
	}{
		{
			title:                   "正常に提案がInsert可能であること",
			preparePersonalContract: nil,
			expectedErr:             nil,
		},
		{
			title: "存在しない提案で登録する場合_エラーが発生すること",
			preparePersonalContract: func(p *personalContractFields) error {
				p.suggestionID = "5332e3d2-2d33-4c63-b622-c85c664c5eab"
				return nil
			},
			expectedErr: xerrors.New(""),
			isLast:      true,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			//ユーザー作成 (メンター募集をするユーザー)
			err := setupUser()
			require.NoError(t, err)
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
			userRepository := NewUserRepositoryImpl(tx)
			err = userRepository.Create(user)
			require.NoError(t, err)

			//メンター募集作成
			err = setupRecruit()
			require.NoError(t, err)
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
			recruitRepository := NewRecruitRepositoryImpl(tx)
			err = recruitRepository.Create(recruit)
			require.NoError(t, err)

			//ユーザー作成 (メンターとなるユーザー)
			err = setupUser()
			require.NoError(t, err)
			user2, err := userdm.NewUser(
				up.userID,
				up.name,
				up.email,
				up.password,
				up.profile,
				userCareers,
				userSkills,
			)
			require.NoError(t, err)
			userRepository2 := NewUserRepositoryImpl(tx)
			err = userRepository2.Create(user2)
			require.NoError(t, err)

			// メンター作成
			err = setupMentor()
			require.NoError(t, err)
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
			require.NoError(t, err)

			mentorRepository := NewMentorRepositoryImpl(tx)
			err = mentorRepository.Create(mentor)
			require.NoError(t, err)

			// 提案作成
			err = setupSuggestion()
			require.NoError(t, err)
			suggestion, err := suggestiondm.NewSuggestion(
				sp.suggestionID,
				sp.mentorID,
				sp.recruitID,
				sp.price,
				sp.suggestionTypeOnce,
				sp.detail,
				sp.suggestionStatusApproval,
			)
			require.NoError(t, err)

			suggestionRepository := NewSuggestionRepositoryImpl(tx)
			err = suggestionRepository.Create(suggestion)
			require.NoError(t, err)

			// 提案契約作成
			err = setupPersonalContact()
			require.NoError(t, err)
			p := personalContractFields{
				personalContractID:     pp.personalContractID.String(),
				suggestionID:           pp.suggestionID.String(),
				personalContractStatus: pp.unapprovedStatus.Uint16(),
			}

			if td.preparePersonalContract != nil {
				if err := td.preparePersonalContract(&p); err != nil {
					t.Fatalf("prepareSuggestion() error = %v", err)
				}
			}

			personalContract, err := personalcontractdm.Reconstruct(
				p.personalContractID,
				p.suggestionID,
				p.personalContractStatus,
			)
			require.NoError(t, err)

			personalContractRepository := NewPersonalContractRepositoryImpl(tx)
			err = personalContractRepository.Create(personalContract)

			if td.expectedErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				selectedPersonalContract, err := personalContractRepository.FetchByID(personalContract.PersonalContractID())
				require.NoError(t, err)
				expected := pp.personalContractID
				actual := selectedPersonalContract.PersonalContractID()
				asserts.Equal(actual, expected)
			}

			if td.isLast {
				tx.Rollback()
			}
		})
	}
}
