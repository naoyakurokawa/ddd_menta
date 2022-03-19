package recruituc

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/naoyakurokawa/ddd_menta/core/domain/recruitdm/mock_recruitdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func TestCreate(t *testing.T) {
	asserts := assert.New(t)
	type fields struct {
		recruitRepository *mock.MockRecruitRepository
	}

	for _, td := range []struct {
		title         string
		userID        string
		inputTitle    string
		budget        uint32
		recruitType   uint16
		detail        string
		recruitStatus uint16
		expectedErr   error
		prepareMock   func(f *fields) error
	}{
		{
			title:         "recruitStatusが「募集終了」の場合_エラーが発生すること",
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusTerminated.Uint16(),
			expectedErr:   xerrors.New("RecruitStatus must be Draft or Published when create"),
			prepareMock:   nil,
		},
		{
			title:         "recruitStatusが「下書き」の場合_エラーが発生しないこと",
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   nil,
			prepareMock: func(f *fields) error {
				f.recruitRepository.EXPECT().Create(gomock.Any()).Return(nil)
				return nil
			},
		},
		{
			title:         "userIDが空の場合_エラーが発生すること",
			userID:        "",
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   xerrors.New("error NewUserIDByVal"),
			prepareMock:   nil,
		},
		{
			title:         "titleが空の場合_エラーが発生すること",
			userID:        rp.userID.String(),
			inputTitle:    "",
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   xerrors.New("title must not be empty"),
			prepareMock:   nil,
		},
		{
			title:         "budgetが1000円未満の場合_エラーが発生すること",
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        999,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   xerrors.New("budget more than ¥1000"),
			prepareMock:   nil,
		},
		{
			title:         "recruitTypeが0の場合_エラーが発生すること",
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   0,
			detail:        rp.detail,
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   xerrors.New("RecruitType must be 1 or 2"),
			prepareMock:   nil,
		},
		{
			title:         "detailが空の場合_エラーが発生すること",
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        "",
			recruitStatus: rp.recruitStatusDraft.Uint16(),
			expectedErr:   xerrors.New("detial must not be empty"),
			prepareMock:   nil,
		},
		{
			title:         "drecruitStatusが0の場合_エラーが発生すること",
			userID:        rp.userID.String(),
			inputTitle:    rp.title,
			budget:        rp.budget,
			recruitType:   rp.recruitTypeOnce.Uint16(),
			detail:        rp.detail,
			recruitStatus: 0,
			expectedErr:   xerrors.New("RecruitStatus must be 1 or 2 or 3"),
			prepareMock:   nil,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				recruitRepository: mock.NewMockRecruitRepository(ctrl),
			}
			if td.prepareMock != nil {
				if err := td.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}
			createRecruitUsecase := NewCreateRecruitUsecase(f.recruitRepository)

			err := createRecruitUsecase.Create(
				td.userID,
				td.inputTitle,
				td.budget,
				td.recruitType,
				td.detail,
				td.recruitStatus,
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
