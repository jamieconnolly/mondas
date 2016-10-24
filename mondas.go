package mondas

import (
	"log"
	"os"
	"path/filepath"
)

var CommandLine = NewApp(filepath.Base(os.Args[0]))

func AddCommand(cmd Command) {
	CommandLine.AddCommand(cmd)
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
