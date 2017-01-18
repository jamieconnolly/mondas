package cli

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kardianos/osext"
)

// App represents the entire command-line interface.
type App struct {
	// Commands is the list of commands.
	Commands Commands
	// ExecutablePrefix is the prefix for all the program executables.
	ExecutablePrefix string
	// HelpCommand is the help command.
	HelpCommand *Command
	// LibexecDir is the directory where program executables are stored.
	LibexecDir string
	// Name is the name of the application.
	Name string
	// Version is the version of the application.
	Version string
	// VersionCommand is the version command.
	VersionCommand *Command

	initialized bool
}

// NewApp creates a new App with some reasonable defaults.
func NewApp(name string, version string) *App {
	exePath, _ := osext.Executable()

	return &App{
		ExecutablePrefix: filepath.Base(exePath) + "-",
		LibexecDir:       filepath.Join(exePath, "../../libexec"),
		Name:             name,
		Version:          version,
	}
}

// AddCommand adds a command to the application.
func (a *App) AddCommand(cmd *Command) {
	a.Commands.Add(cmd)
}

// Init populates the list of commands with the program executables.
// It also sets up the help and version commands.
func (a *App) Init() error {
	if a.initialized {
		return nil
	}

	files, _ := filepath.Glob(filepath.Join(a.LibexecDir, a.ExecutablePrefix+"*"))
	for _, file := range files {
		if _, err := exec.LookPath(file); err == nil {
			a.AddCommand(&Command{
				Name: strings.TrimPrefix(filepath.Base(file), a.ExecutablePrefix),
				Path: file,
			})
		}
	}

	if a.HelpCommand != nil {
		a.AddCommand(a.HelpCommand)
	} else {
		return errors.New("No help command has been set")
	}

	if a.VersionCommand != nil {
		a.AddCommand(a.VersionCommand)
	} else {
		return errors.New("No version command has been set")
	}

	a.initialized = true
	return nil
}

// LookupCommand returns a command matching the given name.
func (a *App) LookupCommand(name string) *Command {
	return a.Commands.Lookup(name)
}

// Run parses the given argument list and runs the matching command.
func (a *App) Run(arguments []string) error {
	if err := a.Init(); err != nil {
		return err
	}

	args := Args(arguments)

	if args.Len() == 0 {
		args = append(args, a.HelpCommand.Name)
	}

	switch args.First() {
	case "--completion":
		return a.ShowCompletions()

	case "--help", "-h":
		args[0] = a.HelpCommand.Name

	case "--version", "-v":
		args[0] = a.VersionCommand.Name
	}

	if cmd := a.LookupCommand(args.First()); cmd != nil && cmd.Runnable() {
		return cmd.Run(NewContext(a, args[1:], os.Environ()))
	}

	return a.ShowInvalidCommandError(args.First())
}

func (a *App) ShowCompletions() error {
	for _, cmd := range a.Commands.Sort() {
		fmt.Println(cmd.Name)
	}
	return nil
}

func (a *App) ShowHelp() error {
	fmt.Printf("Usage: %s <command> [<args>]\n", a.Name)

	if len(a.Commands) > 0 {
		fmt.Println("\nCommands:")

		for _, cmd := range a.Commands.LoadMetadata().Sort() {
			fmt.Printf("   %-15s   %s\n", cmd.Name, cmd.Summary)
		}
	}

	return nil
}

func (a *App) ShowInvalidCommandError(typedName string) error {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "'%s' is not a valid command.\n", typedName)

	if suggestions := a.Commands.SuggestionsFor(typedName); len(suggestions) > 0 {
		if len(suggestions) == 1 {
			fmt.Fprintln(buf, "\nDid you mean this?")
		} else {
			fmt.Fprintln(buf, "\nDid you mean one of these?")
		}
		for _, cmd := range suggestions.Sort() {
			fmt.Fprintf(buf, "\t%s\n", cmd.Name)
		}
	}

	return fmt.Errorf(buf.String())
}
