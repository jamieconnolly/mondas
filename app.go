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
		return fmt.Errorf("Usage: %s <command> [<args>]", a.Name)
	}

	if cmd := a.Find(arguments[0]); cmd != nil {
		return cmd.Run(arguments[1:])
	}

	return fmt.Errorf("'%s' is not a %s command.", arguments[0], a.Name)
}
