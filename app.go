package mondas

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kardianos/osext"
)

type App struct {
	executablePrefix string
	libexecDir string
	name string
}

func NewApp(name string) *App {
	binDir, _ := osext.ExecutableFolder()

	return &App{
		executablePrefix: name + "-",
		libexecDir: filepath.Join(binDir, "..", "libexec"),
		name: name,
	}
}

func (a *App) ExecutablePrefix() string {
	return a.executablePrefix
}

func (a *App) FindAll() Commands {
	commands := Commands{}
	files, _ := filepath.Glob(filepath.Join(a.libexecDir, a.executablePrefix + "*"))
	for _, file := range files {
		if isExecutable(file) {
			cmdName := strings.TrimPrefix(filepath.Base(file), a.executablePrefix)
			commands = append(commands, NewCommand(cmdName, file))
		}
	}
	return commands.Sort()
}

func (a *App) LibexecDir() string {
	return a.libexecDir
}

func (a *App) Lookup(cmdName string) *Command {
	for _, c := range a.FindAll() {
		if c.Name == cmdName {
			return c
		}
	}
	return nil
}

func (a *App) Name() string {
	return a.name
}

func (a *App) Run(arguments []string) error {
	if len(arguments) == 0 {
		return a.ShowHelp()
	}

	switch(arguments[0]) {
	case "--completion":
		return a.ShowCompletions()

	case "help", "--help", "-h":
		if len(arguments) > 1 {
			if cmd := a.Lookup(arguments[1]); cmd != nil {
				return cmd.ShowHelp()
			}
			return a.ShowInvalidCommandError(arguments[1])
		}
		return a.ShowHelp()
	}

	if cmd := a.Lookup(arguments[0]); cmd != nil {
		return cmd.Run(arguments[1:])
	}

	return a.ShowInvalidCommandError(arguments[0])
}

func (a *App) ShowCompletions() error {
	for _, cmd := range a.FindAll() {
		fmt.Println(cmd.Name)
	}
	return nil
}

func (a *App) ShowHelp() error {
	fmt.Printf("Usage: %s <command> [<args>]\n", a.Name())

	if commands := a.FindAll(); len(commands) > 0 {
		fmt.Println("\nCommands:")
		for _, cmd := range commands {
			cmd.Parse()
			fmt.Printf("   %-15s   %s\n", cmd.Name, cmd.Summary)
		}
	}

	return nil
}

func (a *App) ShowInvalidCommandError(typedName string) error {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "'%s' is not a valid command.\n", typedName)

	if suggestions := a.FindAll().SuggestionsFor(typedName); len(suggestions) > 0 {
		if len(suggestions) == 1 {
			fmt.Fprintln(buf, "\nDid you mean this?")
		} else {
			fmt.Fprintln(buf, "\nDid you mean one of these?")
		}
		for _, cmd := range suggestions {
			fmt.Fprintf(buf, "\t%s\n", cmd.Name)
		}
	}

	return fmt.Errorf(buf.String())
}
