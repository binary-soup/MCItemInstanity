package writers

import (
	"item_insanity/cmds/build/builders"
	"item_insanity/cmds/build/data"
	"item_insanity/config"
	"path/filepath"
)

var ADVANCEMENT_PATH = filepath.Join("data", builders.PACK_NAMESPACE, "advancement")

type AdvancementWriter struct{}

func (AdvancementWriter) WriteRoot(cfg *config.Config, info *data.Info, out string) error {
	err := writeJSON("root", cfg.JoinDatapack(ADVANCEMENT_PATH, out), builders.InfoBuilder{}.BuildRoot(info.Display))
	if err != nil {
		return err
	}

	return nil
}
