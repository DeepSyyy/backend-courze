package helper

import "time"

func GetCurrentTime() time.Time {
	return time.Now().Local()
}
