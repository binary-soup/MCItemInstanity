package build_cmd

import (
	"item_insanity/cmds/build/writers"
	"os"

	"github.com/binary-soup/go-command/style"
	"github.com/binary-soup/go-command/util"
)

func (cmd BuildCommand) runClean() error {
	files := []string{writers.MCMETA_FILE, "data"}

	for _, file := range files {
		err := cmd.removeFiles(cmd.cfg.JoinDatapack(file))
		if err != nil {
			return err
		}
	}

	return nil
}

func (BuildCommand) removeFiles(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return util.ChainError(err, "error removing file")
	}

	style.Delete.PrintF("- %s\n", path)
	return nil
}
