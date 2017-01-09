package cli

import (
	"bytes"
	"errors"
	"fmt"
)

type App struct {
	commands       Commands
	helpCommand    Command
	initialized    bool
	name           string
	version        string
	versionCommand Command
}

func NewApp(name string, version string) *App {
	return &App{
		name:    name,
		version: version,
	}
}

func (a *App) AddCommand(cmd Command) {
	a.commands.Add(cmd)
}

func (a *App) HelpCommand() Command {
	return a.helpCommand
}

func (a *App) Init() error {
	if a.initialized {
		return nil
	}

	if a.helpCommand != nil {
		a.AddCommand(a.helpCommand)
	} else {
		return errors.New("No help command has been set")
	}

	if a.versionCommand != nil {
		a.AddCommand(a.versionCommand)
	} else {
		return errors.New("No version command has been set")
	}

	a.initialized = true
	return nil
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

	switch args.First() {
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

func (a *App) SetHelpCommand(cmd Command) {
	a.helpCommand = cmd
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
		for _, cmd := range commands.LoadMetadata() {
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
