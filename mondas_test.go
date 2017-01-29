package mondas_test

import (
	"testing"

	"github.com/jamieconnolly/mondas"
	"github.com/jamieconnolly/mondas/cli"
	// "github.com/jamieconnolly/mondas/commands"
	"github.com/stretchr/testify/assert"
)

func TestAddCommand(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}
	cmds := cli.Commands{cmd1}
	mondas.CommandLine = &cli.App{Commands: cmds}

	mondas.AddCommand(cmd2)
	assert.Equal(t, 2, len(mondas.CommandLine.Commands))
	assert.Equal(t, cmd1.Name, mondas.CommandLine.Commands[0].Name)
	assert.Equal(t, cmd2.Name, mondas.CommandLine.Commands[1].Name)
}
