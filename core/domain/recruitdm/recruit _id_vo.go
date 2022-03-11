package recruitdm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type RecruitID sharedvo.ID

func NewRecruitID() RecruitID {
	return RecruitID(sharedvo.NewID())
}

func NewRecruitIDByVal(strId string) (RecruitID, error) {
	id, err := sharedvo.NewIDByVal(strId)
	if err != nil {
		return RecruitID(""), xerrors.New("error NewRecruitIDByVal")
	}
	return RecruitID(id), nil
}

func NewEmptyRecruitID() RecruitID {
	return RecruitID(sharedvo.NewEmptyID())
}

func (i RecruitID) Equals(i2 RecruitID) bool {
	return i == i2
}

func (i RecruitID) String() string {
	return string(i)
}
