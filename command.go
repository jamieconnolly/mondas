package mondas

import (
	"os"
	"syscall"
)

type Command struct {
	Name string
	Path string
}

func (c *Command) Run(arguments []string) error {
	args := append([]string{c.Path}, arguments...)

	return syscall.Exec(c.Path, args, os.Environ())
}
