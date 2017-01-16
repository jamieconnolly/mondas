package commands

import "github.com/jamieconnolly/mondas/cli"

var HelpCommand = &cli.Command{
	Name:    "help",
	Summary: "Display help information",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args

		if args.Len() > 0 {
			if cmd := ctx.App.LookupCommand(args.First()); cmd != nil && cmd.Runnable() {
				return cmd.ShowHelp()
			}

			return ctx.App.ShowInvalidCommandError(args.First())
		}

		return ctx.App.ShowHelp()
	},
}
