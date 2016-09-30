package mondas

import (
	"log"
	"path/filepath"
)

type App struct {
	LibexecDir string
	Name string
}

func (a *App) Find(command string) *Command {
	return &Command{
		Name: command,
		Path: filepath.Join(a.LibexecDir, filepath.Base(a.Name) + "-" + command),
	}
}

func (a *App) Run(arguments []string) error {
	if len(arguments) == 0 {
		log.Fatalf("Usage: %s <command> [<args>]", a.Name)
	}

	cmd := a.Find(arguments[0])

	return cmd.Run(arguments[1:])
}
