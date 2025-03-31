package extract_cmd

import "regexp"

var ITEM_FILTER = []*regexp.Regexp{
	regexp.MustCompile("^air$"),
	regexp.MustCompile("^banner$"),
}

func (i *Item) Filter() bool {
	for _, filter := range ITEM_FILTER {
		if filter.MatchString(i.ID) {
			return true
		}
	}
	return false
}
