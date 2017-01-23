package cli_test

import (
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	app := cli.NewApp("foo")
	assert.Equal(t, "foo-", app.ExecPrefix)
	assert.Equal(t, "foo", app.Name)
	assert.Equal(t, "foo <command> [<args>]", app.Usage)
}

func TestApp_AddCommand(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}
	cmds := cli.Commands{cmd1}
	app := &cli.App{Commands: cmds}

	app.AddCommand(cmd2)
	assert.Equal(t, 2, len(app.Commands))
	assert.Equal(t, cmd1.Name, app.Commands[0].Name)
	assert.Equal(t, cmd2.Name, app.Commands[1].Name)
}

func TestApp_Initialized(t *testing.T) {
	app := &cli.App{}
	assert.False(t, app.Initialized())

	app.Initialize()
	assert.True(t, app.Initialized())
}

func TestApp_LookupCommand(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}
	cmds := cli.Commands{cmd1, cmd2}

	app := &cli.App{Commands: cmds}
	assert.Equal(t, cmd1.Name, app.LookupCommand("one").Name)
	assert.Equal(t, cmd2.Name, app.LookupCommand("two").Name)
	assert.Nil(t, app.LookupCommand("three"))
}
