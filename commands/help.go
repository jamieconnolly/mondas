package commands

import (
	"fmt"

	"github.com/jamieconnolly/mondas/cli"
)

// ShowAppHelp displays the help information for the given app.
func ShowAppHelp(ctx *cli.Context) error {
	fmt.Printf("Usage: %s %s\n", ctx.App.Name, ctx.App.Usage)

	if len(ctx.App.Commands) > 0 {
		fmt.Println("\nCommands:")

		for _, cmd := range ctx.App.Commands.LoadMetadata(ctx).Sort() {
			fmt.Printf("   %-15s   %s\n", cmd.Name, cmd.Summary)
		}
	}
	return nil
}

// ShowCommandHelp displays the help information for the given command.
func ShowCommandHelp(ctx *cli.Context, c *cli.Command) error {
	c.LoadMetadata(ctx)

	fmt.Println("Name:")
	fmt.Printf("   %s - %s\n", c.Name, c.Summary)

	fmt.Println("\nUsage:")
	fmt.Printf("   %s %s %s\n", ctx.App.Name, c.Name, c.Usage)

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
			return ShowAppHelp(ctx)
		}

		if cmd := ctx.App.LookupCommand(args.First()); cmd != nil {
			return ShowCommandHelp(ctx, cmd)
		}

		return ctx.App.ShowUnknownCommandError(args.First())
	},
}
