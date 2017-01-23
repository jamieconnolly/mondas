package commands

import (
	"fmt"

	"github.com/jamieconnolly/mondas/cli"
)

// VersionCommand displays the version information.
var VersionCommand = &cli.Command{
	Hidden:  true,
	Name:    "version",
	Summary: "Display version information",
	Action: func(ctx *cli.Context) error {
		fmt.Printf("%s version %s\n", ctx.App.Name, ctx.App.Version)
		return nil
	},
}
