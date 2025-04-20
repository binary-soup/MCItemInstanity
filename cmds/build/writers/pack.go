package writers

import (
	"item_insanity/cmds/build/builders"
	"item_insanity/cmds/build/data"
	"item_insanity/config"
)

const (
	MCMETA_FILE = "pack.mcmeta"
	AUTHOR_FILE = "author.json"
	ROOT_FILE   = "root.json"
)

type PackWriter struct{}

func (PackWriter) WriteMeta(cfg *config.Config, pack *data.Pack) error {
	err := writeJSON("mcmeta", cfg.JoinDatapack(MCMETA_FILE), builders.PackBuilder{}.BuildMeta(pack))
	if err != nil {
		return err
	}

	return nil
}

func (PackWriter) WriteRoot(cfg *config.Config, pack *data.Pack) error {
	return AdvancementWriter{}.WriteRoot(cfg, &pack.Root, ROOT_FILE)
}

func (PackWriter) WriteAuthor(cfg *config.Config, pack *data.Pack) error {
	err := writeJSON("author", cfg.JoinDatapack(ADVANCEMENT_PATH, AUTHOR_FILE), builders.PackBuilder{}.BuildAuthor(pack.Author))
	if err != nil {
		return err
	}

	return nil
}
