package personalcontractdm

import (
	"fmt"
	"os"
	"testing"

	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
	"golang.org/x/xerrors"
)

type personalContractParams struct {
	personalContractID     PersonalContractID
	suggestionID           suggestiondm.SuggestionID
	personalContractStatus PersonalContractStatus
	createdAt              sharedvo.CreatedAt
	updatedAt              sharedvo.UpdatedAt
}

var (
	pp personalContractParams
)

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		fmt.Printf("%+v", err)
		return
	}
	os.Exit(m.Run())
}

func setup() error {
	personalContractID := NewPersonalContractID()
	suggestionID := suggestiondm.NewSuggestionID()
	personalContractStatus, err := NewPersonalContractStatus(uint16(1))
	if err != nil {
		return xerrors.New("error NewPersonalContractStatus")
	}
	pp = personalContractParams{
		personalContractID,
		suggestionID,
		personalContractStatus,
		sharedvo.NewCreatedAt(),
		sharedvo.NewUpdatedAt(),
	}
	return nil
}
