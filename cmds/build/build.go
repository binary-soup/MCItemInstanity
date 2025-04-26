package build_cmd

import (
	"item_insanity/cmds/build/builders"
	"item_insanity/cmds/build/data"
	"item_insanity/cmds/build/writers"
	"item_insanity/common"
	"item_insanity/config"
	"path/filepath"

	"github.com/binary-soup/go-command/style"
	"github.com/binary-soup/go-command/util"
)

const (
	PACK_FILE = "pack.json"
)

type buildVisitor struct {
	cfg *config.Config
}

func (v buildVisitor) ParseDirectory(_, dir string) error {
	style.Bolded.Println(dir)
	return nil
}

func (v buildVisitor) ParseRoot(path, dir, file string) error {
	root, err := data.LoadRootJSON(filepath.Join(path, dir, file))
	if err != nil {
		return err
	}

	err = writers.AdvancementWriter{}.WriteRoot(v.cfg, root, filepath.Join(dir, file))
	if err != nil {
		return util.ChainError(err, "error building root advancement")
	}

	return nil
}

func (v buildVisitor) ParseCollect(path, dir, file string) error {
	collect, err := data.LoadCollectJSON(filepath.Join(path, dir, file))
	if err != nil {
		return err
	}

	err = writers.CollectWriter{}.WriteCollect(v.cfg, collect, filepath.Join(dir, file))
	if err != nil {
		return util.ChainError(err, "error building collect advancement")
	}

	return nil
}

func (v buildVisitor) ParseAll(path, dir, file string) error {
	all, err := data.LoadCollectAllJSON(filepath.Join(path, dir, file))
	if err != nil {
		return err
	}

	visitor := common.NewInventoryVisitor()
	parser := common.TreeParser{
		Visitor: &visitor,
	}

	err = parser.Parse(path, dir)
	if err != nil {
		return util.ChainError(err, "error parsing collect tree")
	}

	collect := data.Collect{
		Name:   all.Display.Group,
		Parent: "root",
		Display: data.CollectDisplay{
			Item:        all.Display.Item,
			Title:       all.Display.Title,
			Description: all.Display.Description,
			Frame:       builders.FRAME_CHALLENGE,
		},
		Items: visitor.Ids,
	}

	err = writers.CollectWriter{}.WriteCollect(v.cfg, &collect, filepath.Join(dir, file))
	if err != nil {
		return util.ChainError(err, "error building collect all advancement")
	}

	return nil
}

func (cmd BuildCommand) runBuild() error {
	err := cmd.buildPack()
	if err != nil {
		return err
	}

	err = cmd.buildCollectAdvancements(cmd.cfg.StaticDataPath(), common.ITEM_TREE_ROOT)
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

func (cmd BuildCommand) buildCollectAdvancements(path, dir string) error {
	parser := common.TreeParser{
		Visitor: buildVisitor{
			cfg: cmd.cfg,
		},
	}

	return parser.Parse(path, dir)
}
