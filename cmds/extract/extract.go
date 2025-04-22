package extract_cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/style"
	"github.com/binary-soup/go-command/util"
)

const (
	LANG_PATH      = "assets/minecraft/lang/en_us.json"
	ITEMS_PATH     = "data/items"
	RAW_ITEMS_FILE = "raw.md"
)

func (cmd ExtractCommand) runExtract() error {
	ids, items, err := cmd.extractItems(true)
	if err != nil {
		return util.ChainError(err, "error extracting items")
	}

	err = cmd.writeRawItemTable(ids, items)
	if err != nil {
		return util.ChainError(err, "error writing table")
	}

	return nil
}

func (cmd ExtractCommand) extractItems(verbose bool) ([]string, ItemMap, error) {
	lang, err := os.Open(filepath.Join(cmd.cfg.MinecraftData, LANG_PATH))
	if err != nil {
		return nil, nil, util.ChainError(err, "error opening minecraft lang file")
	}
	defer lang.Close()

	ids := []string{}
	items := ItemMap{}
	duplicates := 0

	scanner := bufio.NewScanner(lang)
	for scanner.Scan() {
		item := Item{}

		if !item.Parse(scanner.Text()) {
			continue
		}

		_, ok := items[item.ID]
		if ok {
			duplicates++
			continue
		}

		items[item.ID] = item
		ids = append(ids, item.ID)
	}

	if verbose {
		style.Create.PrintF("+ %d item(s) extracted\n", len(items))
		style.Delete.PrintF("- %d duplicate(s) filtered\n", duplicates)
	}

	return ids, items, nil
}

func (cmd ExtractCommand) writeRawItemTable(ids []string, items ItemMap) error {
	path := cmd.cfg.JoinRoot(ITEMS_PATH, RAW_ITEMS_FILE)

	file, err := os.Create(path)
	if err != nil {
		return util.ChainError(err, "error creating raw item table file")
	}
	defer file.Close()

	style.BoldCreate.PrintF("+ %s\n", path)

	table := ItemTableWriter{
		File: file,
	}

	table.WriteHeader("Raw Item List")
	table.WriteTable(items, ids)

	fmt.Fprintln(file)
	return nil
}
