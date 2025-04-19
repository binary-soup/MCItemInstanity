package data

import (
	"fmt"

	"github.com/binary-soup/go-command/util"
)

const (
	ROOT_PACK_TITLE = "Item Insanity"
)

func LoadRootJSON(path string) (*Info, error) {
	info, err := util.LoadJSON[Info]("root", path)
	if err != nil {
		return nil, util.ChainError(err, "error loading root json")
	}

	info.Name = "root"
	info.Display.Title = buildRootTitle(info.Display.Title)

	return info, nil
}

func buildRootTitle(title string) string {
	return fmt.Sprintf("%s [%s]", title, ROOT_PACK_TITLE)
}
