package util

import "time"

const (
	TimeForMat = "2006-01-02"
)

func ParseTimeStr(timeStr, formatStr string) (time.Time, error) {
	return time.Parse(formatStr, timeStr)
}
