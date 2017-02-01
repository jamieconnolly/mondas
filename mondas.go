package mondas

import (
	"os"
	"path/filepath"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/jamieconnolly/mondas/commands"
	"github.com/kardianos/osext"
)

var exePath, _ = osext.Executable()

// These variables are used to configure the default application.
var (
	// Commands is the list of commands
	Commands cli.Commands
	// ExecPath is the directory where program executables are stored
	ExecPath = filepath.Join(exePath, "../../libexec")
	// Name is the name of the application
	Name = filepath.Base(os.Args[0])
	// Version is the version of the application
	Version string
)

// AddCommand adds a command to the default application.
func AddCommand(cmd *cli.Command) {
	Commands.Add(cmd)
}

// Run is the main entry point for the command-line interface.
// It creates the default application, adds the helper commands,
// and then runs it using the arguments from os.Args.
func Run() {
	app := cli.NewApp(Name)
	app.Commands = Commands
	app.ExecPath = ExecPath
	app.Version = Version

	app.AddCommand(commands.CompletionsCommand)
	app.AddCommand(commands.HelpCommand)
	app.AddCommand(commands.VersionCommand)

	cli.Exit(app.Run(os.Args[1:]))
}
