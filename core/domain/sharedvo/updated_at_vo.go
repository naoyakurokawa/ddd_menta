package sharedvo

import "time"

type UpdatedAt time.Time

func NewUpdatedAt() UpdatedAt {
	now := time.Now()
	return UpdatedAt(now)
}

func (ua UpdatedAt) Time() time.Time {
	return time.Time(ua)
}
