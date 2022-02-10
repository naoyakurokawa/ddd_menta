package sharedvo

import "time"

type CreatedAt time.Time

func NewCreatedAt() CreatedAt {
	now := time.Now()
	return CreatedAt(now)
}
