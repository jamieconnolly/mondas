package mondas

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kardianos/osext"
)

type App struct {
	commands Commands
	executablePrefix string
	helpCommand Command
	initialized bool
	libexecDir string
	name string
	version string
	versionCommand Command
}

func NewApp(name string) *App {
	binDir, _ := osext.ExecutableFolder()

	return &App{
		executablePrefix: name + "-",
		helpCommand: helpCommand,
		libexecDir: filepath.Join(binDir, "..", "libexec"),
		name: name,
		versionCommand: versionCommand,
	}
}

func (a *App) AddCommand(cmd Command) {
	a.commands.Add(cmd)
}

func (a *App) Commands() Commands {
	return a.commands
}

func (a *App) ExecutablePrefix() string {
	return a.executablePrefix
}

func (a *App) HelpCommand() Command {
	return a.helpCommand
}

func (a *App) Init() *App {
	if a.initialized {
		return a
	}

	files, _ := filepath.Glob(filepath.Join(a.libexecDir, a.executablePrefix + "*"))
	for _, file := range files {
		if isExecutable(file) {
			name := strings.TrimPrefix(filepath.Base(file), a.executablePrefix)
			a.commands.Add(NewExecCommand(name, file))
		}
	}

	a.AddCommand(a.helpCommand)

	if a.version != "" {
		a.AddCommand(a.versionCommand)
	}

	a.initialized = true
	return a
}

func (a *App) LibexecDir() string {
	return a.libexecDir
}

func (a *App) LookupCommand(name string) Command {
	return a.commands.Lookup(name)
}

func (a *App) Name() string {
	return a.name
}

func (a *App) Run(arguments []string) error {
	args := Args(arguments)

	if args.Len() == 0 {
		args = append(args, a.helpCommand.Name())
	}

	switch(args.First()) {
	case "--completion":
		return a.ShowCompletions()

	case "--help", "-h":
		args[0] = a.helpCommand.Name()

	case "--version", "-v":
		args[0] = a.versionCommand.Name()
	}

	if cmd := a.LookupCommand(args.First()); cmd != nil {
		return cmd.Run(NewContext(a, Args(args[1:])))
	}

	return a.ShowInvalidCommandError(args.First())
}

func (a *App) SetExecutablePrefix(prefix string) {
	a.executablePrefix = prefix
}

func (a *App) SetHelpCommand(cmd Command) {
	a.helpCommand = cmd
}

func (a *App) SetLibexecDir(dir string) {
	a.libexecDir = dir
}

func (a *App) SetName(name string) {
	a.name = name
}

func (a *App) SetVersion(version string) {
	a.version = version
}

func (a *App) SetVersionCommand(cmd Command) {
	a.versionCommand = cmd
}

func (a *App) ShowCompletions() error {
	for _, cmd := range a.commands.Sort() {
		fmt.Println(cmd.Name())
	}
	return nil
}

func (a *App) ShowHelp() error {
	fmt.Printf("Usage: %s <command> [<args>]\n", a.Name())

	if commands := a.commands.Sort(); len(commands) > 0 {
		fmt.Println("\nCommands:")
		for _, cmd := range commands {
			cmd.LoadHelp()
			fmt.Printf("   %-15s   %s\n", cmd.Name(), cmd.Summary())
		}
	}

	return nil
}

func (a *App) ShowInvalidCommandError(typedName string) error {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "'%s' is not a valid command.\n", typedName)

	if suggestions := a.commands.SuggestionsFor(typedName); len(suggestions) > 0 {
		if len(suggestions) == 1 {
			fmt.Fprintln(buf, "\nDid you mean this?")
		} else {
			fmt.Fprintln(buf, "\nDid you mean one of these?")
		}
		for _, cmd := range suggestions {
			fmt.Fprintf(buf, "\t%s\n", cmd.Name())
		}
	}

	return fmt.Errorf(buf.String())
}

func (a *App) Version() string {
	return a.version
}

func (a *App) VersionCommand() Command {
	return a.versionCommand
}
