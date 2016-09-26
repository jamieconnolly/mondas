package mondas

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func Run() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <command> [<args>]\n", os.Args[0])
		os.Exit(1)
	}

	binDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	cmdName := filepath.Base(os.Args[0]) + "-" + os.Args[1]
	cmdPath := filepath.Join(binDir, "../libexec", cmdName)

	args := []string{cmdPath}
	args = append(args, os.Args[2:]...)

	if err := syscall.Exec(cmdPath, args, os.Environ()); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", cmdName, err)
	}
}
