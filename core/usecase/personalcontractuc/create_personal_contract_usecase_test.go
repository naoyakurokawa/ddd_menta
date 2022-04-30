package personalcontractuc

import (
	"testing"

	"github.com/golang/mock/gomock"
	personalContractMock "github.com/naoyakurokawa/ddd_menta/core/domain/personalcontractdm/mock_personalcontractdm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
	suggestionMock "github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm/mock_suggestiondm"
	"github.com/naoyakurokawa/ddd_menta/customerrors"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func TestCreate(t *testing.T) {
	type fields struct {
		personalContractRepository *personalContractMock.MockPersonalContractRepository
		suggestionRepository       *suggestionMock.MockSuggestionRepository
	}

	for _, td := range []struct {
		title        string
		suggestionID string
		expectedErr  error
		prepareMock  func(f *fields) error
	}{
		{
			title:        "SuggestionStatusがUnapprovedのとき_エラーが発生しないこと",
			suggestionID: pp.suggestionID.String(),
			expectedErr:  nil,
			prepareMock: func(f *fields) error {
				suggestion, err := suggestiondm.Reconstruct(
					sp.suggestionID.String(),
					sp.mentorID.String(),
					sp.recruitID.String(),
					sp.price,
					sp.suggestionTypeOnce.Uint16(),
					sp.detail,
					sp.suggestionStatusUnapproved.Uint16(),
				)
				if err != nil {
					return err
				}

				f.personalContractRepository.EXPECT().Create(gomock.Any()).Return(nil)
				f.suggestionRepository.EXPECT().FetchByID(gomock.Any()).Return(suggestion, nil)

				return nil
			},
		},
		{
			title:        "SuggestionStatusがUnapprovedではない場合_エラーが発生すること",
			suggestionID: pp.suggestionID.String(),
			expectedErr:  customerrors.NewInvalidParameter(),
			prepareMock: func(f *fields) error {
				suggestion, err := suggestiondm.Reconstruct(
					sp.suggestionID.String(),
					sp.mentorID.String(),
					sp.recruitID.String(),
					sp.price,
					sp.suggestionTypeOnce.Uint16(),
					sp.detail,
					sp.suggestionStatusTerminated.Uint16(),
				)
				if err != nil {
					return err
				}
				f.suggestionRepository.EXPECT().FetchByID(gomock.Any()).Return(suggestion, nil)

				return nil
			},
		},
		{
			title:        "SuggestionIDが空の場合_エラーが発生すること",
			suggestionID: "",
			expectedErr:  xerrors.New("error NewSuggestionIDByVal"),
			prepareMock:  nil,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				personalContractRepository: personalContractMock.NewMockPersonalContractRepository(ctrl),
				suggestionRepository:       suggestionMock.NewMockSuggestionRepository(ctrl),
			}
			if td.prepareMock != nil {
				if err := td.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}
			personalContractUsecase := NewCreatePersonalContractUsecase(f.personalContractRepository, f.suggestionRepository)
			err := personalContractUsecase.Create(
				td.suggestionID,
			)

			if td.expectedErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
