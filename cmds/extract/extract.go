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

const LANG_PATH = "assets/minecraft/lang/en_us.json"
const RAW_ITEM_LIST = "data/items/raw_item_list.txt"

type ExtractCommand struct {
	command.CommandBase
}

func NewExtractCommand() ExtractCommand {
	return ExtractCommand{
		CommandBase: command.NewCommandBase("extract", "extract all item ids from Minecraft's data files"),
	}
}

func (cmd ExtractCommand) Run(args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return util.ChainError(err, "error loading config")
	}

	input := filepath.Join(cfg.MinecraftData, LANG_PATH)
	output := cfg.JoinRoot(RAW_ITEM_LIST)

	os.MkdirAll(filepath.Dir(output), 0700)

	err = extractIds(input, output)
	if err != nil {
		return util.ChainError(err, "error extracting ids")
	}

	return nil
}

func extractIds(input, output string) error {
	inFile, err := os.Open(input)
	if err != nil {
		return util.ChainError(err, "error opening input file")
	}
	defer inFile.Close()

	outFile, err := os.Create(output)
	if err != nil {
		return util.ChainError(err, "error creating output file")
	}
	defer outFile.Close()

	style.BoldCreate.PrintF("+ %s\n", output)

	count := 0
	filtered := 0

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		item := Item{}

		if !item.Parse(scanner.Text()) {
			continue
		}

		if item.Filter() {
			filtered++
			continue
		}

		count++
		item.Write(outFile)
	}

	style.Create.PrintF("Count: %d\n", count)
	style.Delete.PrintF("Filtered: %d\n", filtered)

	return nil
}
