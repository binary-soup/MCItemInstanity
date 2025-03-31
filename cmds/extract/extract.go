package extract_cmd

import (
	"fmt"
	"item_insanity/config"

	"github.com/binary-soup/go-command/command"
	"github.com/binary-soup/go-command/util"
)

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

	fmt.Println(cfg.MinecraftDataPath)
	return nil
}
