package extract_cmd

import (
	"item_insanity/config"

	"github.com/binary-soup/go-command/command"
	"github.com/binary-soup/go-command/util"
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
	diff := cmd.Flags.Bool("diff", false, "calculate the diff between the extracted items and the static files")
	cmd.Flags.Parse(args)

	var err error

	cmd.cfg, err = config.Load()
	if err != nil {
		return util.ChainError(err, "error loading config")
	}

	if *diff {
		return cmd.runDiff()
	} else {
		return cmd.runExtract()
	}
}
