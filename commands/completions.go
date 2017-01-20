package commands

import (
	"fmt"

	"github.com/jamieconnolly/mondas/cli"
)

// CompletionsCommand displays the list of commands for autocompletion.
var CompletionsCommand = &cli.Command{
	Name:    "completions",
	Summary: "Display the list of commands for autocompletion",
	Action: func(ctx *cli.Context) error {
		for _, cmd := range ctx.App.Commands.Sort() {
			fmt.Println(cmd.Name)
		}
		return nil
	},
}
