package build_cmd

import (
	"item_insanity/cmds/build/data"
	"item_insanity/cmds/build/writers"
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/style"
	"github.com/binary-soup/go-command/util"
)

const (
	PACK_FILE = "pack.json"
	ROOT_FILE = "root.json"
)

func (cmd BuildCommand) runBuild() error {
	err := cmd.buildPack()
	if err != nil {
		return err
	}

	err = cmd.buildCollectAdvancements("", cmd.cfg.StaticDataPath())
	if err != nil {
		return err
	}

	return nil
}

func (cmd BuildCommand) buildPack() error {
	pack, err := data.LoadPackJSON(cmd.cfg.JoinStaticData(PACK_FILE))
	if err != nil {
		return err
	}

	style.Bolded.Println("pack")

	err = writers.PackWriter{}.WriteMeta(cmd.cfg, pack)
	if err != nil {
		return util.ChainError(err, "error building mcmeta")
	}

	writer := writers.PackWriter{}

	err = writer.WriteRoot(cmd.cfg, pack)
	if err != nil {
		return util.ChainError(err, "error building pack root advancement")
	}

	err = writer.WriteAuthor(cmd.cfg, pack)
	if err != nil {
		return util.ChainError(err, "error building pack author advancement")
	}

	return nil
}

func (cmd BuildCommand) buildCollectAdvancements(dir, path string) error {
	if dir != "" {
		style.Bolded.PrintF("advancements/%s\n", dir)
	}

	entires, err := os.ReadDir(path)
	if err != nil {
		return util.ChainError(err, "error reading directory")
	}

	subDirs := []string{}

	for _, entry := range entires {
		if entry.IsDir() {
			subDirs = append(subDirs, entry.Name())
			continue
		}

		inFile := filepath.Join(path, entry.Name())
		outFile := filepath.Join(dir, entry.Name())

		switch entry.Name() {
		case PACK_FILE:
			continue
		case ROOT_FILE:
			err = cmd.buildRootAdvancement(inFile, outFile)
		default:
			err = cmd.buildCollectAdvancement(inFile, outFile)
		}

		if err != nil {
			return err
		}
	}

	for _, subDir := range subDirs {
		err = cmd.buildCollectAdvancements(filepath.Join(dir, subDir), filepath.Join(path, subDir))
		if err != nil {
			return err
		}
	}

	return nil
}

func (cmd BuildCommand) buildRootAdvancement(src, out string) error {
	root, err := data.LoadRootJSON(src)
	if err != nil {
		return err
	}

	err = writers.AdvancementWriter{}.WriteRoot(cmd.cfg, root, out)
	if err != nil {
		return util.ChainError(err, "error building root advancement")
	}

	return nil
}

func (cmd BuildCommand) buildCollectAdvancement(src, out string) error {
	collect, err := data.LoadCollectJSON(src)
	if err != nil {
		return err
	}

	err = writers.CollectWriter{}.WriteCollect(cmd.cfg, collect, out)
	if err != nil {
		return util.ChainError(err, "error building collect advancement")
	}

	return nil
}
