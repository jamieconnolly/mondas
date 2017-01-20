package mondas

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/jamieconnolly/mondas/commands"
	"github.com/kardianos/osext"
)

// CommandLine is the default application.
var CommandLine = cli.NewApp(filepath.Base(os.Args[0]))

// Version is the default version of the application.
var Version string

// AddCommand adds a command to the default application.
func AddCommand(cmd *cli.Command) {
	CommandLine.AddCommand(cmd)
}

// Run runs the default application using the arguments from os.Args.
func Run() {
	log.SetFlags(0)
	log.SetPrefix(CommandLine.Name + ": ")

	if exePath, err := osext.Executable(); err == nil {
		CommandLine.ExecPath = filepath.Join(exePath, "../../libexec")
	}

	CommandLine.HelpCommand = commands.HelpCommand
	CommandLine.Version = Version

	CommandLine.AddCommand(commands.CompletionsCommand)

	if Version != "" {
		CommandLine.AddCommand(commands.VersionCommand)
	}

	if err := CommandLine.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
