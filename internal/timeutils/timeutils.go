package timeutils

import "time"

const goReferenceTime = "20060102T150405"

func GenerateCurrentDateTimeString() string {
	now := time.Now()

	return now.Format(goReferenceTime)
}
