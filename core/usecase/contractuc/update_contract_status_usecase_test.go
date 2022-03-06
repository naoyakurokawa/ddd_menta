package contractuc

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naoyakurokawa/ddd_menta/core/domain/contractdm"
	mock "github.com/naoyakurokawa/ddd_menta/core/domain/contractdm/mock_contractdm"
	"github.com/stretchr/testify/require"
)

func TestUpdateContractStatus(t *testing.T) {
	type fields struct {
		contractRepository *mock.MockContractRepository
	}

	for _, td := range []struct {
		title                 string
		contractID            string
		requestContractStatus contractdm.ContractStatus
		expectedErr           error
		prepareMock           func(f *fields) error
	}{
		{
			title:                 "Contractが未承認で_契約中に更新しようとするとき_エラーが発生しないこと",
			contractID:            cp.contractID.String(),
			requestContractStatus: cp.underContractStatus,
			expectedErr:           nil,
			prepareMock: func(f *fields) error {
				contract, err := contractdm.NewContract(
					cp.contractID,
					up.userID,
					mp.mentorID,
					mp.planID,
					cp.unapprovedStatus,
				)
				if err != nil {
					return err
				}

				f.contractRepository.EXPECT().FindByID(gomock.Any()).Return(contract, nil)
				f.contractRepository.EXPECT().UpdateContractStatus(gomock.Any(), gomock.Any()).Return(nil)

				return nil
			},
		},
		{
			title:                 "Contractが未承認で_契約終了に更新しようとするとき_エラーが発生すること",
			contractID:            cp.contractID.String(),
			requestContractStatus: cp.terminatedContractStatus,
			expectedErr:           errors.New("can't update contract"),
			prepareMock: func(f *fields) error {
				contract, err := contractdm.NewContract(
					cp.contractID,
					up.userID,
					mp.mentorID,
					mp.planID,
					cp.unapprovedStatus,
				)
				if err != nil {
					return err
				}

				f.contractRepository.EXPECT().FindByID(gomock.Any()).Return(contract, nil)

				return nil
			},
		},
		{
			title:                 "リクエストするContractIDが空の場合_エラーが発生すること",
			contractID:            "",
			requestContractStatus: 0,
			expectedErr:           errors.New("error NewContractIDByVal"),
			prepareMock:           nil,
		},
		{
			title:                 "リクエストするContractStatusが1-3の範囲外の場合_エラーが発生すること",
			contractID:            cp.contractID.String(),
			requestContractStatus: 0,
			expectedErr:           errors.New("ContractStatus must be 1 or 2 or 3"),
			prepareMock:           nil,
		},
	} {
		t.Run(td.title, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				contractRepository: mock.NewMockContractRepository(ctrl),
			}
			if td.prepareMock != nil {
				if err := td.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}
			contractUsecase := NewUpdateContractStatusUsecase(f.contractRepository)
			err := contractUsecase.UpdateContractStatus(
				td.contractID,
				td.requestContractStatus.Uint16(),
			)

			if td.expectedErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
