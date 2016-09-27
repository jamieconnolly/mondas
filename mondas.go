package mondas

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

func Run() {
	log.SetFlags(0)

	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flags.SetOutput(ioutil.Discard)

	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalf("%s: %v", os.Args[0], err)
	}

	if flags.NArg() == 0 {
		log.Fatalf("Usage: %s <command> [<args>]", os.Args[0])
	}

	binDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	cmdName := filepath.Base(os.Args[0]) + "-" + flags.Arg(0)
	cmdPath := filepath.Join(binDir, "../libexec", cmdName)

	args := []string{cmdPath}
	args = append(args, flags.Args()[1:]...)

	if err := syscall.Exec(cmdPath, args, os.Environ()); err != nil {
		log.Fatalf("%s: %v", cmdName, err)
	}
}
