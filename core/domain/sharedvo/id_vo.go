package sharedvo

import (
	"regexp"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type ID string

var (
	IDFormat = `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`
	IDRegExp = regexp.MustCompile(IDFormat)
)

func NewID() ID {
	return ID(uuid.New().String())
}

func NewIDByVal(id string) (ID, error) {
	if ok := IDRegExp.MatchString(id); !ok {
		return ID(""), xerrors.Errorf("invalid ID format. ID is %s", id)
	}
	return ID(id), nil
}

func NewEmptyID() ID {
	return ID("")
}

func (i ID) String() string {
	return string(i)
}

func (i ID) Equals(i2 ID) bool {
	return i == i2
}
