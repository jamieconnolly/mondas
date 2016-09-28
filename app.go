package mondas

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

type App struct {
	Name string
}

func NewApp(name string) *App {
	return &App{
		Name: name,
	}
}

func (a *App) Find(command string) string {
	binDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	cmdName := filepath.Base(a.Name) + "-" + command
	cmdPath := filepath.Join(binDir, "../libexec", cmdName)

	return cmdPath
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

	cmdPath := a.Find(flags.Arg(0))

	args := []string{cmdPath}
	args = append(args, flags.Args()[1:]...)

	if err := syscall.Exec(cmdPath, args, os.Environ()); err != nil {
		log.Fatalf("%s-%s: %v", filepath.Base(a.Name), flags.Arg(0), err)
	}
}
