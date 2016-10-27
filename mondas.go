package mondas

import (
	"log"
	"os"
	"path/filepath"
)

var (
	CommandLine = New(Name, Version)
	Name string
	Version string
)

func AddCommand(cmd Command) {
	CommandLine.AddCommand(cmd)
}

func New(name string, version string) *App {
	if name == "" {
		name = filepath.Base(os.Args[0])
	}

	if version == "" {
		version = "0.0.0"
	}

	return NewApp(name, version)
}

func Run() {
	log.SetFlags(0)
	log.SetPrefix(CommandLine.Name() + ": ")

	if err := CommandLine.Init().Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func SetHelpCommand(cmd Command) {
	CommandLine.SetHelpCommand(cmd)
}

func SetVersionCommand(cmd Command) {
	CommandLine.SetVersionCommand(cmd)
}
