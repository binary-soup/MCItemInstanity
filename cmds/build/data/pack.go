package data

import "github.com/binary-soup/go-command/util"

type Pack struct {
	Description string `json:"description"`
	Root        Info   `json:"root"`
}

func LoadPackJSON(path string) (*Pack, error) {
	pack, err := util.LoadJSON[Pack]("pack", path)
	if err != nil {
		return nil, util.ChainError(err, "error loading pack json")
	}

	pack.Root.Name = "root"
	pack.Root.Display.Title = buildRootTitle(pack.Root.Display.Title)

	return pack, nil
}
