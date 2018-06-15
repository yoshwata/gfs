package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/yoshwata/gfs/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "ngram",
		Usage:  "",
		Action: command.CmdNgram,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "comp",
		Usage:  "",
		Action: command.CmdComp,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "gendata",
		Usage:  "",
		Action: command.CmdGenData,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
