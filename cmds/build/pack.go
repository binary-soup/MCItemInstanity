package build_cmd

import (
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/style"
	"github.com/binary-soup/go-command/util"
)

const PACK_PATH = "data/static/pack.json"
const MC_META_PATH = "data/datapack/pack.mcmeta"

func (cmd BuildCommand) buildPack() error {
	pack, err := util.LoadJSON[pack]("pack", cmd.config.JoinRoot(PACK_PATH))
	if err != nil {
		return util.ChainError(err, "error loading pack file")
	}

	err = cmd.buildMCMeta(pack)
	if err != nil {
		return err
	}

	return nil
}

func (cmd BuildCommand) buildMCMeta(pack *pack) error {
	meta := mcMeta{
		Pack: mcMetaPack{
			Description: pack.Description,
			PackFormat:  48,
		},
	}

	path := cmd.config.JoinRoot(MC_META_PATH)
	os.MkdirAll(filepath.Dir(path), 0700)

	err := writeJSON("mc meta", path, meta)
	if err != nil {
		return util.ChainError(err, "error building mc meta")
	}

	style.BoldCreate.PrintF("+ %s\n", path)
	return nil
}
