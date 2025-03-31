package extract_cmd

import (
	"bufio"
	"item_insanity/config"
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/command"
	"github.com/binary-soup/go-command/util"
)

const LANG_PATH = "assets/minecraft/lang/en_us.json"
const OUT_FILE = "raw_item_list.txt"

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

	input := filepath.Join(cfg.MinecraftDataPath, LANG_PATH)
	output := cfg.JoinRoot(OUT_FILE)

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

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		item := Item{}

		ok := item.Parse(scanner.Text())
		if !ok {
			continue
		}

		item.Write(outFile)
	}

	return nil
}
