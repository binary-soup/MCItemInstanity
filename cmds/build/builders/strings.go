package builders

import "strings"

func idToLowerSpaced(id string) string {
	return strings.ReplaceAll(id, "_", " ")
}

func idToUpperSpaced(id string) string {
	return strings.ToUpper(idToLowerSpaced(id))
}
