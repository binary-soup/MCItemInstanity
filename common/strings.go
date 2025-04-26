package common

import "strings"

func ToLowerSpaced(str string) string {
	return strings.ReplaceAll(str, "_", " ")
}

func ToUpperSpaced(str string) string {
	return strings.ToUpper(ToLowerSpaced(str))
}

func Pluralize(count int, single, plural string) string {
	if count == 1 {
		return single
	} else {
		return plural
	}
}
