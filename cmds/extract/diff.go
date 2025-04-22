package extract_cmd

import (
	"os"

	"github.com/binary-soup/go-command/style"
	"github.com/binary-soup/go-command/util"
)

const (
	EXCLUDE_FILE    = "exclude.json"
	DIFF_ITEMS_FILE = "diff.md"
)

type Exclude struct {
	Items []string `json:"items"`
}

func (cmd ExtractCommand) runDiff() error {
	// TODO: load item set from the tree

	raw, items, ids, err := cmd.loadRawSet()
	if err != nil {
		return util.ChainError(err, "error loading raw set")
	}

	exclude, err := cmd.loadExcludeSet()
	if err != nil {
		return util.ChainError(err, "error loading exclude set")
	}

	diff := cmd.calcDiff(raw, exclude)

	err = cmd.writeDiffTables(items, ids, diff, exclude)
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

	style.Info.PrintF("+ %d item(s) not (yet) included\n", len(set))

	return set, items, ids, nil
}

func (cmd ExtractCommand) loadExcludeSet() (ItemIdSet, error) {
	exclude, err := util.LoadJSON[Exclude]("exclude", cmd.cfg.JoinStaticData(EXCLUDE_FILE))
	if err != nil {
		return nil, err
	}

	set := make(ItemIdSet, len(exclude.Items))
	for _, id := range exclude.Items {
		set[id] = struct{}{}
	}

	style.Delete.PrintF("- %d item(s) excluded\n", len(set))

	return set, nil
}

func (cmd ExtractCommand) calcDiff(diff, compare ItemIdSet) ItemIdSet {
	set := make(ItemIdSet, len(diff)-len(compare))

	for id := range diff {
		if _, ok := compare[id]; !ok {
			set[id] = struct{}{}
		}
	}

	return set
}

func (cmd ExtractCommand) writeDiffTables(items ItemMap, ids []string, diff, exclude ItemIdSet) error {
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
	cmd.writeItemSet(table, items, ids, exclude)

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
