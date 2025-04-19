package writers

import (
	"item_insanity/cmds/build/builders"
	"item_insanity/cmds/build/data"
	"item_insanity/config"
)

const (
	MCMETA_FILE = "pack.mcmeta"
)

type PackWriter struct{}

func (PackWriter) WriteMeta(cfg *config.Config, pack *data.Pack) error {
	builder := builders.PackBuilder{}

	err := writeJSON("mcmeta", cfg.JoinDatapack(MCMETA_FILE), builder.BuildMeta(pack))
	if err != nil {
		return err
	}

	return nil
}
