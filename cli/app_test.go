package cli_test

import (
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	app := cli.NewApp("foo")
	assert.Equal(t, "foo", app.Name)
	assert.Equal(t, "foo-", app.ExecutablePrefix)
}

func TestApp_AddCommand(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}
	cmds := cli.Commands{cmd1}
	app := &cli.App{Commands: cmds}

	app.AddCommand(cmd2)
	assert.Equal(t, 2, len(app.Commands))
	assert.Equal(t, cmd1, app.Commands[0])
	assert.Equal(t, cmd2, app.Commands[1])
}

func TestApp_LookupCommand(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}
	cmds := cli.Commands{cmd1, cmd2}

	app := &cli.App{Commands: cmds}
	assert.Equal(t, cmd1, app.LookupCommand("one"))
	assert.Equal(t, cmd2, app.LookupCommand("two"))
	assert.Nil(t, app.LookupCommand("three"))
}
