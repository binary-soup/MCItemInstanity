package main

import (
	"flag"
	"fmt"
	"os"

	extract_cmd "item_insanity/cmds/extract"

	"github.com/binary-soup/go-command/command"
	"github.com/binary-soup/go-command/style"
)

func main() {
	ls := flag.Bool("ls", false, "list all commands")
	flag.Parse()

	runner := command.NewRunner(extract_cmd.NewExtractCommand())

	if *ls || len(os.Args) < 2 {
		runner.ListCommands()
		return
	}

	if err := runner.RunCommand(os.Args[1], os.Args[2:]); err != nil {
		style.BoldError.Print("ERROR: ")
		fmt.Println(err)
	}
}
