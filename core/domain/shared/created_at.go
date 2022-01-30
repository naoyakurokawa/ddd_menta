package shared

import "time"

type CreatedAt time.Time

func GetCurrentTime() CreatedAt {
	now := time.Now()
	return CreatedAt(now)
}
