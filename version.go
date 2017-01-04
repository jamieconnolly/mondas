package mondas

import "fmt"

type VersionCommand struct {
	name    string
	summary string
}

var versionCommand = &VersionCommand{
	name:    "version",
	summary: "Display version information",
}

func (c *VersionCommand) LoadHelp() error {
	return nil
}

func (c *VersionCommand) Name() string {
	return c.name
}

func (c *VersionCommand) Run(ctx *Context) error {
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
