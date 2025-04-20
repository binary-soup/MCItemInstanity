package writers

import (
	"item_insanity/cmds/build/builders"
	"item_insanity/cmds/build/data"
	"item_insanity/config"
	"path/filepath"
)

type CollectWriter struct{}

func (CollectWriter) WriteCollect(cfg *config.Config, collect *data.Collect, out string) error {
	err := writeJSON("collect", cfg.JoinDatapack(ADVANCEMENT_PATH, out), builders.CollectBuilder{}.BuildCollect(filepath.Dir(out), collect))
	if err != nil {
		return err
	}

	return nil
}
