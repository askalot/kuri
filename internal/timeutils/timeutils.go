package timeutils

import "time"

func GenerateCurrentDateTimeString() string {
	const goReferenceTime = "20060102T150405"

	now := time.Now().UTC()

	return now.Format(goReferenceTime)
}
