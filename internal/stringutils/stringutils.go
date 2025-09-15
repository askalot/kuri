package timeutils

import (
	"strings"
)

func ConvertCSVStringToSlice(values string) []string {
	var list []string

	for _, value := range strings.Split(values, ",") {
		list = append(list, value)
	}

	return list
}
