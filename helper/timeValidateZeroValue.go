package helper

import (
	"time"
)

func TimeParseing() time.Time {
	time, err := time.Parse("2006-01-02 15:04:05-07:00", "0001-01-01 00:00:00-00:00")
	if err != nil {
		panic(err)
	}

	return time
}
