package mondas

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/jamieconnolly/mondas/commands"
)

var (
	// CommandLine is the default application.
	CommandLine = cli.NewApp(Name, Version)
	// HelpCommand is the default help command.
	HelpCommand = commands.HelpCommand
	// Name is the default name of the application.
	Name = filepath.Base(os.Args[0])
	// Version is the default version of the application.
	Version = ""
	// VersionCommand is the default version command.
	VersionCommand = commands.VersionCommand
)

// AddCommand adds a command to the default application.
func AddCommand(cmd *cli.Command) {
	CommandLine.AddCommand(cmd)
}

// Run runs the default application using the arguments from os.Args.
func Run() {
	log.SetFlags(0)
	log.SetPrefix(CommandLine.Name + ": ")

	CommandLine.HelpCommand = HelpCommand
	CommandLine.VersionCommand = VersionCommand

	if err := CommandLine.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
