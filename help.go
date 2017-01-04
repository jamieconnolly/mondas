package mondas

import "fmt"

type HelpCommand struct {
	name    string
	summary string
}

var helpCommand = &HelpCommand{
	name:    "help",
	summary: "Display help information",
}

func (c *HelpCommand) LoadHelp() error {
	return nil
}

func (c *HelpCommand) Name() string {
	return c.name
}

func (c *HelpCommand) Run(ctx *Context) error {
	args := ctx.Args

	if args.Len() > 0 {
		for _, arg := range args {
			switch arg {
			case "--help", "-h":
				return c.ShowHelp()
			}
		}

		if cmd := ctx.App.LookupCommand(args.First()); cmd != nil {
			return cmd.ShowHelp()
		}

		return ctx.App.ShowInvalidCommandError(args.First())
	}

	return ctx.App.ShowHelp()
}

func (c *HelpCommand) ShowHelp() error {
	fmt.Println("Name:")
	fmt.Printf("   %s - %s\n", c.Name(), c.Summary())

	if len(c.Usage()) > 0 {
		fmt.Println("\nUsage:")
		for _, use := range c.Usage() {
			fmt.Printf("   %s\n", use)
		}
	}

	return nil
}

func (c *HelpCommand) Summary() string {
	return c.summary
}

func (c *HelpCommand) Usage() []string {
	return nil
}
