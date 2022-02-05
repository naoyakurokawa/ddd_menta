package sharedvo

import (
	"github.com/google/uuid"
)

type ID string

func NewID() ID {
	return ID(uuid.New().String())
}

func NewIDByVal(id string) ID {
	return ID(id)
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
