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

		for _, c := range a.Commands.Parse().Sort() {
			fmt.Printf("   %-15s   %s\n", c.Name, c.Summary)
		}
	}
	return nil
}

// ShowCommandHelp displays the help information for the given command.
func ShowCommandHelp(c *cli.Command) error {
	c.Parse()

	fmt.Println("Name:")
	fmt.Printf("   %s - %s\n", c.Name, c.Summary)

	fmt.Println("\nUsage:")
	fmt.Printf("   %s\n", c.Usage)

	return nil
}

// HelpCommand displays the help information.
var HelpCommand = &cli.Command{
	Name:    "help",
	Summary: "Display help information",
	Usage:   "<command>",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args

		if args.Len() == 0 {
			return ShowAppHelp(ctx.App)
		}

		if cmd := ctx.App.LookupCommand(args.First()); cmd != nil {
			return ShowCommandHelp(cmd)
		}

		return ctx.App.ShowUnknownCommandError(args.First())
	},
}
