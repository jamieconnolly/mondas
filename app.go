package mondas

import (
	"fmt"
	"path/filepath"
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

func (a *App) Run(arguments []string) error {
	if len(arguments) == 0 {
		return a.ShowHelp()
	}

	if cmd := a.Find(arguments[0]); cmd != nil {
		return cmd.Run(arguments[1:])
	}

	return a.ShowInvalidCommandHelp(arguments[0])
}

func (a *App) ShowHelp() error {
	fmt.Printf("Usage: %s <command> [<args>]\n", a.Name)
	return nil
}

func (a *App) ShowInvalidCommandHelp(typedCommand string) error {
	return fmt.Errorf("%s: '%s' is not a %s command.", a.Name, typedCommand, a.Name)
}
