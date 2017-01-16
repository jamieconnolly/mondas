package mondas

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/jamieconnolly/mondas/commands"
	"github.com/kardianos/osext"
)

var (
	CommandLine      = cli.NewApp(Name)
	Executable, _    = osext.Executable()
	ExecutablePrefix = filepath.Base(Executable) + "-"
	LibexecDir       = filepath.Join(Executable, "../../libexec")
	Name             = filepath.Base(Executable)
	Version          = ""
)

func AddCommand(cmd *cli.Command) {
	CommandLine.AddCommand(cmd)
}

func Run() {
	log.SetFlags(0)
	log.SetPrefix(CommandLine.Name + ": ")

	CommandLine.ExecutablePrefix = ExecutablePrefix
	CommandLine.HelpCommand = commands.HelpCommand
	CommandLine.LibexecDir = LibexecDir
	CommandLine.Version = Version
	CommandLine.VersionCommand = commands.VersionCommand

	if err := CommandLine.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
