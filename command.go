package mondas

import (
	"os"
	"syscall"
)

type Command struct {
	Name string
	Path string
}

func NewCommand(name, path string) *Command {
	return &Command{
		Name: name,
		Path: path,
	}
}

func (c *Command) Run(arguments []string) error {
	args := []string{c.Path}
	args = append(args, arguments...)

	return syscall.Exec(c.Path, args, os.Environ())
}
