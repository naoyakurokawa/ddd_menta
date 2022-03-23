package repoimpl

import (
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

type suggestionFields struct {
	suggestionID     string
	mentorID         string
	recruitID        string
	price            uint32
	suggestionType   uint16
	detail           string
	suggestionStatus uint16
}

func TestSuggestionRepoCreate(t *testing.T) {
	tx := db.NewDB().Begin()
	asserts := assert.New(t)
	for _, td := range []struct {
		title             string
		prepareSuggestion func(s *suggestionFields) error
		expectedErr       error
		isLast            bool
	}{
		{
			title:             "正常に提案がInsert可能であること",
			prepareSuggestion: nil,
			expectedErr:       nil,
		},
		{
			title: "存在しないメンターで登録する場合_エラーが発生すること",
			prepareSuggestion: func(s *suggestionFields) error {
				s.mentorID = "5332e3d2-2d33-4c63-b622-c85c664c5eab"
				return nil
			},
			expectedErr: xerrors.New(""),
		},
		{
			title: "存在しない募集に提案する場合_エラーが発生すること",
			prepareSuggestion: func(s *suggestionFields) error {
				s.recruitID = "2332e3d2-2d33-4c63-b622-c85c664c5eab"
				return nil
			},
			expectedErr: xerrors.New(""),
			isLast:      true,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			//ユーザー作成 (メンター募集をするユーザー)
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
			userRepository := NewUserRepositoryImpl(tx)
			err = userRepository.Create(user)
			require.NoError(t, err)

			//メンター募集作成
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
			recruitRepository := NewRecruitRepositoryImpl(tx)
			err = recruitRepository.Create(recruit)
			require.NoError(t, err)

			//ユーザー作成 (メンターとなるユーザー)
			setupUser()
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
			setupMentor()
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
			setupSuggestion()
			s := suggestionFields{
				suggestionID:     sp.suggestionID.String(),
				mentorID:         sp.mentorID.String(),
				recruitID:        sp.recruitID.String(),
				price:            sp.price,
				suggestionType:   sp.suggestionTypeOnce.Uint16(),
				detail:           sp.detail,
				suggestionStatus: sp.suggestionStatusApproval.Uint16(),
			}

			if td.prepareSuggestion != nil {
				if err := td.prepareSuggestion(&s); err != nil {
					t.Fatalf("prepareSuggestion() error = %v", err)
				}
			}

			suggestion, err := suggestiondm.Reconstruct(
				s.suggestionID,
				s.mentorID,
				s.recruitID,
				s.price,
				s.suggestionType,
				s.detail,
				s.suggestionStatus,
			)
			require.NoError(t, err)

			suggestionRepository := NewSuggestionRepositoryImpl(tx)
			err = suggestionRepository.Create(suggestion)

			if td.expectedErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				selectedSuggestion, err := suggestionRepository.FetchByID(suggestion.SuggestionID())
				require.NoError(t, err)
				expected := sp.suggestionID
				actual := selectedSuggestion.SuggestionID()
				asserts.Equal(actual, expected)
			}

			if td.isLast {
				tx.Rollback()
			}
		})
	}
}
