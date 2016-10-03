package mondas

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
)

type App struct {
	LibexecDir string
	Name string
}

func (a *App) Find(command string) *Command {
	file := filepath.Join(a.LibexecDir, a.Name + "-" + command)
	if isExecutable(file) {
		return &Command{
			Name: command,
			Path: file,
		}
	}
	return nil
}

func (a *App) FindAll() []*Command {
	commands := []*Command{}
	files, _ := filepath.Glob(filepath.Join(a.LibexecDir, a.Name + "-*"))
	for _, file := range files {
		if isExecutable(file) {
			commands = append(commands, &Command{
				Name: strings.TrimPrefix(filepath.Base(file), a.Name + "-"),
				Path: file,
			})
		}
	}
	return commands
}

func (a *App) FindSuggestionsFor(typedCommand string) []*Command {
	suggestions := []*Command{}
	for _, cmd := range a.FindAll() {
		suggestForDistance := stringDistance(typedCommand, cmd.Name) <= MaxSuggestionDistance
		suggestForPrefix := strings.HasPrefix(strings.ToLower(cmd.Name), strings.ToLower(typedCommand))
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

	if cmd := a.Find(arguments[0]); cmd != nil {
		return cmd.Run(arguments[1:])
	}

	return a.ShowInvalidCommandError(arguments[0])
}

func (a *App) ShowHelp() error {
	fmt.Printf("Usage: %s <command> [<args>]\n", a.Name)
	return nil
}

func (a *App) ShowInvalidCommandError(typedCommand string) error {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s: '%s' is not a valid command.\n", a.Name, typedCommand)
	if suggestions := a.FindSuggestionsFor(typedCommand); len(suggestions) > 0 {
		fmt.Fprintln(buf)
		if len(suggestions) == 1 {
			fmt.Fprintln(buf, "Did you mean this?")
		} else {
			fmt.Fprintln(buf, "Did you mean one of these?")
		}
		for _, s := range suggestions {
			fmt.Fprintf(buf, "\t%v\n", s.Name)
		}
	}
	return fmt.Errorf(buf.String())
}
