package data

import (
	"path/filepath"
	"strings"

	"github.com/binary-soup/go-command/util"
)

type Collect struct {
	Name    string
	Parent  string         `json:"parent"`
	Display CollectDisplay `json:"display"`
	Type    string         `json:"type"`
	Items   []string       `json:"items"`
}

type CollectDisplay struct {
	Item  string `json:"item"`
	Title string `json:"title"`
	Frame string `json:"frame"`
}

func LoadCollectJSON(path string) (*Collect, error) {
	c, err := util.LoadJSON[Collect]("collect", path)
	if err != nil {
		return nil, util.ChainError(err, "error loading collect json")
	}

	ext := filepath.Ext(path)
	c.Name, _ = strings.CutSuffix(filepath.Base(path), ext)

	return c, nil
}
