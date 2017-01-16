package commands

import (
	"fmt"

	"github.com/jamieconnolly/mondas/cli"
)

var VersionCommand = &cli.Command{
	Name:    "version",
	Summary: "Display version information",
	Action: func(ctx *cli.Context) error {
		if ctx.App.Version != "" {
			fmt.Printf("%s version %s\n", ctx.App.Name, ctx.App.Version)
		}
		return nil
	},
}
