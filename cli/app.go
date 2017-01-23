package cli

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// App represents the entire command-line interface.
type App struct {
	// Commands is the list of commands
	Commands Commands
	// ExecPath is the directory where program executables are stored
	ExecPath string
	// ExecPrefix is the prefix for program executables
	ExecPrefix string
	// HelpCommand is the help command.
	HelpCommand *Command
	// Name is the name of the application
	Name string
	// Usage is the one-line usage message
	Usage string
	// Version is the version of the application
	Version string

	initialized bool
}

// NewApp creates a new App with some reasonable defaults.
func NewApp(name string) *App {
	return &App{
		ExecPrefix: name + "-",
		Name:       name,
		Usage:      name + " <command> [<args>]",
	}
}

// AddCommand adds a command to the application.
func (a *App) AddCommand(cmd *Command) {
	a.Commands.Add(cmd)
}

// Init prepends the exec path to PATH then populates the list
// of commands with program executables and the help command.
func (a *App) Init() error {
	if a.initialized {
		return nil
	}

	if a.ExecPath != "" {
		os.Setenv("PATH", strings.Join(
			[]string{a.ExecPath, os.Getenv("PATH")},
			string(os.PathListSeparator),
		))
	}

	for _, dir := range filepath.SplitList(os.Getenv("PATH")) {
		if dir == "" {
			dir = "."
		}
		files, _ := filepath.Glob(filepath.Join(dir, a.ExecPrefix+"*"))
		for _, file := range files {
			if _, err := exec.LookPath(file); err == nil {
				a.AddCommand(&Command{
					Name: strings.TrimPrefix(filepath.Base(file), a.ExecPrefix),
					Path: file,
				})
			}
		}
	}

	if a.HelpCommand != nil {
		a.AddCommand(a.HelpCommand)
	} else {
		return errors.New("No help command has been set")
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

	args[0] = strings.TrimPrefix(args.First(), "--")

	if args.Contains("--help") {
		args = Args([]string{a.HelpCommand.Name, args.First()})
	}

	if cmd := a.LookupCommand(args.First()); cmd != nil {
		return cmd.Run(NewContext(a, args[1:], os.Environ()))
	}

	return a.ShowUnknownCommandError(args.First())
}

// ShowUnknownCommandError shows a list of suggested commands
// based on the given name then exits with status code 1.
func (a *App) ShowUnknownCommandError(typedName string) error {
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
