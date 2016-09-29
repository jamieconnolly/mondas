package mondas

import (
	"log"
	"os"
	"path/filepath"
)

type App struct {
	Name string
}

func (a *App) Find(command string) *Command {
	binDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	cmdName := filepath.Base(a.Name) + "-" + command
	cmdPath := filepath.Join(binDir, "../libexec", cmdName)

	return &Command{
		Name: command,
		Path: cmdPath,
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
