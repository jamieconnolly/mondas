package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	app := NewApp("foo", "1.2.3")
	assert.Equal(t, "foo", app.Name)
	assert.Equal(t, "foo-", app.ExecutablePrefix)
	assert.Equal(t, "1.2.3", app.Version)
}

func TestApp_AddCommand(t *testing.T) {
	cmd1 := &Command{Name: "one"}
	cmd2 := &Command{Name: "two"}
	cmds := Commands{cmd1}
	app := &App{Commands: cmds}

	app.AddCommand(cmd2)
	assert.Equal(t, 2, len(app.Commands))
	assert.Equal(t, cmd1, app.Commands[0])
	assert.Equal(t, cmd2, app.Commands[1])
}

func TestApp_LookupCommand(t *testing.T) {
	cmd1 := &Command{Name: "one"}
	cmd2 := &Command{Name: "two"}
	cmds := Commands{cmd1, cmd2}

	app := &App{Commands: cmds}
	assert.Equal(t, cmd1, app.LookupCommand("one"))
	assert.Equal(t, cmd2, app.LookupCommand("two"))
	assert.Nil(t, app.LookupCommand("three"))
}
