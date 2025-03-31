package build_cmd

import (
	"item_insanity/config"

	"github.com/binary-soup/go-command/command"
	"github.com/binary-soup/go-command/util"
)

type BuildCommand struct {
	command.CommandBase
	config *config.Config
}

func NewBuildCommand() BuildCommand {
	return BuildCommand{
		CommandBase: command.NewCommandBase("build", "build the datapack"),
	}
}

func (cmd BuildCommand) Run(args []string) error {
	var err error

	cmd.config, err = config.Load()
	if err != nil {
		return util.ChainError(err, "error loading config")
	}

	err = cmd.buildPack()
	if err != nil {
		return util.ChainError(err, "error building datapack pack files")
	}

	return nil
}
