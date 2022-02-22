package util

import "time"

func IsUrlExpired(createDate time.Time) bool {
	const EXPIRED_DURATION = 365 * 24 * time.Hour
	now := time.Now()
	// if expired => return true
	return now.Sub(createDate) > EXPIRED_DURATION
}
