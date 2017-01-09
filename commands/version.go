package commands

import (
	"fmt"

	"github.com/jamieconnolly/mondas/cli"
)

type VersionCommand struct {
	name string
	summary string
}

func NewVersionCommand() *VersionCommand {
	return &VersionCommand{
		name: "version",
		summary: "Display version information",
	}
}

func (c *VersionCommand) LoadMetadata() error {
	return nil
}

func (c *VersionCommand) Name() string {
	return c.name
}

func (c *VersionCommand) Run(ctx *cli.Context) error {
	fmt.Printf("%s version %s\n", ctx.App.Name(), ctx.App.Version())
	return nil
}

func (c *VersionCommand) ShowHelp() error {
	return nil
}

func (c *VersionCommand) Summary() string {
	return c.summary
}

func (c *VersionCommand) Usage() []string {
	return nil
}
