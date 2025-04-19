package build_cmd

import (
	"item_insanity/config"

	"github.com/binary-soup/go-command/command"
	"github.com/binary-soup/go-command/util"
)

type BuildCommand struct {
	command.CommandBase
	cfg *config.Config
}

func NewBuildCommand() BuildCommand {
	return BuildCommand{
		CommandBase: command.NewCommandBase("build", "build the datapack"),
	}
}

func (cmd BuildCommand) Run(args []string) error {
	clean := cmd.Flags.Bool("clean", false, "remove all the generated build files")
	rebuild := cmd.Flags.Bool("rebuild", false, "clean the datapack directory, then build")
	cmd.Flags.Parse(args)

	var err error

	cmd.cfg, err = config.Load()
	if err != nil {
		return util.ChainError(err, "error loading config")
	}

	if *clean || *rebuild {
		err = cmd.runClean()
	}
	if err != nil || *clean {
		return err
	}

	return cmd.runBuild()
}
