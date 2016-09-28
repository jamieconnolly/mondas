package mondas

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type App struct {
	Name string
}

func NewApp(name string) *App {
	return &App{
		Name: name,
	}
}

func (a *App) Find(command string) *Command {
	binDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	cmdName := filepath.Base(a.Name) + "-" + command
	cmdPath := filepath.Join(binDir, "../libexec", cmdName)

	return NewCommand(command, cmdPath)
}

func (a *App) Run(arguments []string) {
	log.SetFlags(0)

	flags := flag.NewFlagSet(a.Name, flag.ContinueOnError)
	flags.SetOutput(ioutil.Discard)

	if err := flags.Parse(arguments); err != nil {
		log.Fatalf("%s: %v", a.Name, err)
	}

	if flags.NArg() == 0 {
		log.Fatalf("Usage: %s <command> [<args>]", a.Name)
	}

	cmd := a.Find(flags.Arg(0))
	if err := cmd.Run(flags.Args()[1:]); err != nil {
		log.Fatalf("%s: %v", a.Name, err)
	}
}
