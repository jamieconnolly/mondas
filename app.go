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

func (a *App) Run(arguments []string) {
	if len(arguments) == 0 {
		log.Fatalf("Usage: %s <command> [<args>]", a.Name)
	}

	cmd := a.Find(arguments[0])

	if err := cmd.Run(arguments[1:]); err != nil {
		log.Fatalf("%s: %v", a.Name, err)
	}
}
