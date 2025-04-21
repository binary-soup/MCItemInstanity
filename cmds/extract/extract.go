package extract_cmd

import (
	"bufio"
	"item_insanity/config"
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/command"
	"github.com/binary-soup/go-command/style"
	"github.com/binary-soup/go-command/util"
)

const (
	LANG_PATH  = "assets/minecraft/lang/en_us.json"
	ITEMS_PATH = "data/items"
)

type ExtractCommand struct {
	command.CommandBase
	cfg *config.Config
}

func NewExtractCommand() ExtractCommand {
	return ExtractCommand{
		CommandBase: command.NewCommandBase("extract", "extract all item ids from Minecraft's data files"),
	}
}

func (cmd ExtractCommand) Run(args []string) error {
	var err error

	cmd.cfg, err = config.Load()
	if err != nil {
		return util.ChainError(err, "error loading config")
	}

	items, err := cmd.extractItems()
	if err != nil {
		return util.ChainError(err, "error extracting ids")
	}

	err = cmd.writeTable(items)
	if err != nil {
		return util.ChainError(err, "error writing table")
	}

	return nil
}

func (cmd ExtractCommand) extractItems() ([]Item, error) {
	lang, err := os.Open(filepath.Join(cmd.cfg.MinecraftData, LANG_PATH))
	if err != nil {
		return nil, util.ChainError(err, "error opening minecraft lang file")
	}
	defer lang.Close()

	items := []Item{}
	duplicates := map[string]uint8{}
	filtered := 0

	scanner := bufio.NewScanner(lang)
	for scanner.Scan() {
		item := Item{}

		if !item.Parse(scanner.Text()) {
			continue
		}

		count, ok := duplicates[item.ID]
		if ok {
			duplicates[item.ID] = count + 1
			filtered++
			continue
		}

		duplicates[item.ID] = 1
		items = append(items, item)
	}

	style.Info.PrintF("+ %d items extracted\n", len(items))
	style.Delete.PrintF("- %d duplicates filtered\n", filtered)

	return items, nil
}

func (cmd ExtractCommand) writeTable(items []Item) error {
	table := ItemTableWriter{}

	err := table.OpenFile(cmd.cfg.JoinRoot(ITEMS_PATH))
	if err != nil {
		return err
	}
	defer table.CloseFile()

	for _, item := range items {
		table.WriteItem(item)
	}

	return nil
}
