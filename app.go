package mondas

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kardianos/osext"
)

type App struct {
	ExecutablePrefix string
	LibexecDir string
	MaxSuggestionDistance int
	Name string
}

func NewApp(name string) *App {
	binDir, _ := osext.ExecutableFolder()

	return &App{
		ExecutablePrefix: name + "-",
		LibexecDir: filepath.Join(binDir, "..", "libexec"),
		MaxSuggestionDistance: 3,
		Name: name,
	}
}

func (a *App) Find(cmdName string) *Command {
	file := filepath.Join(a.LibexecDir, a.ExecutablePrefix + cmdName)
	if isExecutable(file) {
		return NewCommand(cmdName, file)
	}
	return nil
}

func (a *App) FindAll() []*Command {
	commands := []*Command{}
	files, _ := filepath.Glob(filepath.Join(a.LibexecDir, a.ExecutablePrefix + "*"))
	for _, file := range files {
		if isExecutable(file) {
			cmdName := strings.TrimPrefix(filepath.Base(file), a.ExecutablePrefix)
			commands = append(commands, NewCommand(cmdName, file))
		}
	}
	return commands
}

func (a *App) FindSuggested(typedName string) []*Command {
	suggestions := []*Command{}
	for _, cmd := range a.FindAll() {
		suggestForDistance := stringDistance(typedName, cmd.Name) <= a.MaxSuggestionDistance
		suggestForPrefix := strings.HasPrefix(strings.ToLower(cmd.Name), strings.ToLower(typedName))
		if suggestForDistance || suggestForPrefix {
			suggestions = append(suggestions, cmd)
		}
	}
	return suggestions
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
			if cmd := a.Find(arguments[1]); cmd != nil {
				return cmd.ShowHelp()
			}
			return a.ShowInvalidCommandError(arguments[1])
		}
		return a.ShowHelp()
	}

	if cmd := a.Find(arguments[0]); cmd != nil {
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
	fmt.Printf("Usage: %s <command> [<args>]\n", a.Name)

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

	if suggestions := a.FindSuggested(typedName); len(suggestions) > 0 {
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
