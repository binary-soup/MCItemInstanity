package writers

import (
	"item_insanity/cmds/build/builders"
	"item_insanity/cmds/build/data"
	"item_insanity/config"
	"path/filepath"
)

var ADVANCEMENT_PATH = filepath.Join("data", builders.PACK_NAMESPACE, "advancement")

type AdvancementWriter struct{}

func (AdvancementWriter) WriteInfo(cfg *config.Config, info *data.Info, out string) error {
	err := writeJSON("info", cfg.JoinDatapack(ADVANCEMENT_PATH, out), builders.InfoBuilder{}.Build(info))
	if err != nil {
		return err
	}

	return nil
}

func (AdvancementWriter) WriteCollect(cfg *config.Config, collect *data.Collect, out string) error {
	err := writeJSON("collect", cfg.JoinDatapack(ADVANCEMENT_PATH, out), builders.CollectBuilder{}.BuildCollect(filepath.Dir(out), collect))
	if err != nil {
		return err
	}

	return nil
}
