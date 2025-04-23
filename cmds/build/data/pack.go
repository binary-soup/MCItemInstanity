package data

import "github.com/binary-soup/go-command/util"

type Pack struct {
	Description string     `json:"description"`
	Version     string     `json:"version"`
	Root        Info       `json:"root"`
	Author      PackAuthor `json:"author"`
}

type PackAuthor struct {
	Creator string             `json:"creator"`
	Socials []PackAuthorSocial `json:"socials"`
}

type PackAuthorSocial struct {
	Handle string `json:"handle"`
	Color  string `json:"color"`
}

func LoadPackJSON(path string) (*Pack, error) {
	pack, err := util.LoadJSON[Pack]("pack", path)
	if err != nil {
		return nil, util.ChainError(err, "error loading pack json")
	}

	pack.Root.Name = "root"
	//pack.Root.Display.Title = buildRootTitle(pack.Root.Display.Title)

	return pack, nil
}
