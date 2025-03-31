package extract_cmd

import (
	"fmt"
	"io"
	"regexp"
)

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

func (i *Item) Write(w io.Writer) {
	fmt.Fprintf(w, "minecraft:%-40s| [%s] \"%s\"\n", i.ID, i.Type, i.Name)
}
