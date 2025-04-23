package data

import "github.com/binary-soup/go-command/util"

type CollectAll struct {
	Parent  string            `json:"parent"`
	Display CollectAllDisplay `json:"display"`
}

type CollectAllDisplay struct {
	Group       string `json:"group"`
	Item        string `json:"item"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func LoadCollectAllJSON(path string) (*CollectAll, error) {
	all, err := util.LoadJSON[CollectAll]("all", path)
	if err != nil {
		return nil, util.ChainError(err, "error loading collect all json")
	}

	return all, nil
}
