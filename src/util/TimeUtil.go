package util

import (
	"time"
)

const TIMELAYOUT = "2006-01-02 15:04:05"

func TimeToStamp(time time.Time) (timeStamp int64) {
	timeStamp = time.Unix()
	return timeStamp
}

func StampToTime(timeStamp int64) time.Time {
	dateTime := time.Unix(timeStamp, 0)
	return dateTime
}
