package sharedvo

import "time"

type CreatedAt time.Time

func NewCreatedAt() CreatedAt {
	now := time.Now()
	return CreatedAt(now)
}

func (ca CreatedAt) Time() time.Time {
	return time.Time(ca)
}
