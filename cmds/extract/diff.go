package extract_cmd

import (
	"item_insanity/common"
	"os"

	"github.com/binary-soup/go-command/style"
	"github.com/binary-soup/go-command/util"
)

const (
	EXCLUDE_FILE    = "exclude.json"
	DIFF_ITEMS_FILE = "diff.md"
)

type Exclude struct {
	Items map[string][]string `json:"items"`
}

func (cmd ExtractCommand) runDiff() error {
	raw, items, ids, err := cmd.loadRawSet()
	if err != nil {
		return util.ChainError(err, "error loading raw set")
	}

	included, err := cmd.loadIncludedSet()
	if err != nil {
		return util.ChainError(err, "error loading included set")
	}

	excluded, err := cmd.loadExcludedSet()
	if err != nil {
		return util.ChainError(err, "error loading excluded set")
	}

	diff := cmd.calcDiff(raw, included, excluded)
	style.Info.PrintF("x %d %s not (yet) included\n", len(diff), common.Pluralize(len(diff), "item", "items"))

	err = cmd.writeDiffTables(items, ids, diff, excluded, included)
	if err != nil {
		return util.ChainError(err, "error writing diff tables")
	}

	return nil
}

func (cmd ExtractCommand) loadRawSet() (ItemIdSet, ItemMap, []string, error) {
	ids, items, err := cmd.extractItems(false)
	if err != nil {
		return nil, nil, nil, util.ChainError(err, "error extracting items")
	}

	set := make(ItemIdSet, len(ids))
	for _, id := range ids {
		set[id] = struct{}{}
	}

	return set, items, ids, nil
}

func (cmd ExtractCommand) loadExcludedSet() (ItemIdSet, error) {
	exclude, err := util.LoadJSON[Exclude]("exclude", cmd.cfg.JoinStaticData(EXCLUDE_FILE))
	if err != nil {
		return nil, err
	}

	set := ItemIdSet{}

	for key, ids := range exclude.Items {
		for _, id := range ids {
			set[id] = struct{}{}
		}

		style.Delete.PrintF("- %d %s %s\n", len(ids), common.Pluralize(len(ids), "item", "items"), common.ToLowerSpaced(key))
	}

	return set, nil
}

func (cmd ExtractCommand) loadIncludedSet() (ItemIdSet, error) {
	visitor := common.NewInventoryVisitor()
	parser := common.TreeParser{
		Visitor: &visitor,
	}

	err := parser.Parse(cmd.cfg.StaticDataPath(), common.ITEM_TREE_ROOT)
	if err != nil {
		return nil, util.ChainError(err, "error parsing collect tree")
	}

	set := make(ItemIdSet, len(visitor.Ids))
	for _, id := range visitor.Ids {
		set[id] = struct{}{}
	}

	style.Create.PrintF("+ %d %s included\n", len(set), common.Pluralize(len(set), "item", "items"))

	return set, nil
}

func (cmd ExtractCommand) calcDiff(base ItemIdSet, compare ...ItemIdSet) ItemIdSet {
	diff := ItemIdSet{}

	for id := range base {
		if !cmd.anySetContainsId(id, compare...) {
			diff[id] = struct{}{}
		}
	}

	return diff
}

func (cmd ExtractCommand) anySetContainsId(id string, sets ...ItemIdSet) bool {
	for _, set := range sets {
		if _, ok := set[id]; ok {
			return true
		}
	}
	return false
}

func (cmd ExtractCommand) writeDiffTables(items ItemMap, ids []string, diff, excluded, included ItemIdSet) error {
	path := cmd.cfg.JoinRoot(ITEMS_PATH, DIFF_ITEMS_FILE)

	file, err := os.Create(path)
	if err != nil {
		return util.ChainError(err, "error creating diff item tables file")
	}
	defer file.Close()

	style.BoldCreate.PrintF("+ %s\n", path)

	table := ItemTableWriter{
		File: file,
	}

	table.WriteHeader("Not (yet) Included")
	cmd.writeItemSet(table, items, ids, diff)

	table.WriteHeader("Excluded")
	cmd.writeItemSet(table, items, ids, excluded)

	table.WriteHeader("Included")
	cmd.writeItemSet(table, items, ids, included)

	return nil
}

func (cmd ExtractCommand) writeItemSet(table ItemTableWriter, items ItemMap, ids []string, set ItemIdSet) {
	table.WriteTableHeader()

	for _, id := range ids {
		if _, ok := set[id]; ok {
			table.WriteItem(items[id])
		}
	}
}
