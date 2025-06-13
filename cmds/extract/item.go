package extract_cmd

import (
	"regexp"
)

type ItemIdSet map[string]struct{}
type ItemMap map[string]Item

var ITEM_REGEX = regexp.MustCompile(`^\s*"(block|item)\.minecraft\.(\w+)":\s+"(.+)",?$`)

type Item struct {
	Type string
	ID   string
	Name string
}

func (i *Item) Parse(raw string) bool {
	matches := ITEM_REGEX.FindStringSubmatch(raw)
	if len(matches) <= 0 {
		return false
	}

	i.Type = matches[1]
	i.ID = matches[2]
	i.Name = matches[3]

	return true
}
