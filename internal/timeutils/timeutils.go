package timeutils

import "time"

const goReferenceTime = "20060102T150405"

func GenerateCurrentDateTimeString() string {
	return time.Now().Format(goReferenceTime)
}
