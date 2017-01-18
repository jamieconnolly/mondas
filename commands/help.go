package commands

import (
	"fmt"

	"github.com/jamieconnolly/mondas/cli"
)

// ShowAppHelp displays the help information for the given app.
func ShowAppHelp(a *cli.App) error {
	fmt.Printf("Usage: %s\n", a.Usage)

	if len(a.Commands) > 0 {
		fmt.Println("\nCommands:")

		for _, cmd := range a.Commands.LoadMetadata().Sort() {
			fmt.Printf("   %-15s   %s\n", cmd.Name, cmd.Summary)
		}
	}

	return nil
}

// ShowCommandHelp displays the help information for the given command.
func ShowCommandHelp(c *cli.Command) error {
	c.LoadMetadata()

	fmt.Println("Name:")
	fmt.Printf("   %s - %s\n", c.Name, c.Summary)

	if c.Usage != "" {
		fmt.Println("\nUsage:")
		fmt.Printf("   %s\n", c.Usage)
	}

	return nil
}

// HelpCommand is the default help command.
var HelpCommand = &cli.Command{
	Name:    "help",
	Summary: "Display help information",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args

		if args.Len() > 0 {
			if cmd := ctx.App.LookupCommand(args.First()); cmd != nil && cmd.Runnable() {
				return ShowCommandHelp(cmd)
			}

			return ctx.App.ShowInvalidCommandError(args.First())
		}

		return ShowAppHelp(ctx.App)
	},
}
