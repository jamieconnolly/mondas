package mondas

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kardianos/osext"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/jamieconnolly/mondas/commands"
	"github.com/jamieconnolly/mondas/utils"
)

var (
	CommandLine      = cli.NewApp(Name, Version)
	Executable, _    = osext.Executable()
	ExecutablePrefix = filepath.Base(Executable) + "-"
	LibexecDir       = filepath.Join(Executable, "../../libexec")
	Name             = filepath.Base(Executable)
	Version          = ""
)

func AddCommand(cmd cli.Command) {
	CommandLine.AddCommand(cmd)
}

func Run() {
	log.SetFlags(0)
	log.SetPrefix(CommandLine.Name() + ": ")

	for cmdName, cmdFile := range utils.FindAvailableCommands(ExecutablePrefix, LibexecDir) {
		CommandLine.AddCommand(commands.NewExecCommand(cmdName, cmdFile))
	}

	CommandLine.SetHelpCommand(commands.NewHelpCommand())
	CommandLine.SetVersionCommand(commands.NewVersionCommand())

	if err := CommandLine.Init(); err != nil {
		log.Fatal(err)
	}

	if err := CommandLine.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func SetHelpCommand(cmd cli.Command) {
	CommandLine.SetHelpCommand(cmd)
}

func SetName(name string) {
	CommandLine.SetName(name)
}

func SetVersion(version string) {
	CommandLine.SetVersion(version)
}

func SetVersionCommand(cmd cli.Command) {
	CommandLine.SetVersionCommand(cmd)
}
