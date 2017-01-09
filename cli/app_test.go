package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	app := NewApp("foo", "1.2.3")
	assert.Equal(t, "foo", app.Name())
	assert.Equal(t, "1.2.3", app.Version())
}

func TestApp_AddCommand(t *testing.T) {
	cmd1 := &FakeCommand{name: "one"}
	cmd2 := &FakeCommand{name: "two"}
	cmds := Commands{cmd1}
	app := &App{commands: cmds}

	app.AddCommand(cmd2)
	assert.Equal(t, 2, len(app.commands))
	assert.Equal(t, cmd1, app.commands[0])
	assert.Equal(t, cmd2, app.commands[1])
}

func TestApp_LookupCommand(t *testing.T) {
	cmd1 := &FakeCommand{name: "one"}
	cmd2 := &FakeCommand{name: "two"}
	cmds := Commands{cmd1, cmd2}

	app := &App{commands: cmds}
	assert.Equal(t, cmd1, app.LookupCommand("one"))
	assert.Equal(t, cmd2, app.LookupCommand("two"))
	assert.Nil(t, app.LookupCommand("three"))
}

func TestApp_SetHelpCommand(t *testing.T) {
	cmd := &FakeCommand{}
	app := NewApp("foo", "1.2.3")

	app.SetHelpCommand(cmd)
	assert.Equal(t, cmd, app.HelpCommand())
}

func TestApp_SetName(t *testing.T) {
	app := NewApp("foo", "1.2.3")
	app.SetName("bar")
	assert.Equal(t, "bar", app.Name())
}

func TestApp_SetVersion(t *testing.T) {
	app := NewApp("foo", "1.2.3")
	app.SetVersion("4.5.6")
	assert.Equal(t, "4.5.6", app.Version())
}

func TestApp_SetVersionCommand(t *testing.T) {
	cmd := &FakeCommand{}
	app := NewApp("foo", "1.2.3")

	app.SetVersionCommand(cmd)
	assert.Equal(t, cmd, app.VersionCommand())
}
