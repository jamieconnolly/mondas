package cli_test

import (
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/stretchr/testify/assert"
)

func TestCommand_Parse_WithExecutable(t *testing.T) {
	cmd := &cli.Command{Path: "testdata/foo-hello"}
	assert.False(t, cmd.Parsed())

	err := cmd.Parse()
	if assert.NoError(t, err) {
		assert.Equal(t, "Display \"Hello, world!\"", cmd.Summary)
		assert.Equal(t, "foo hello", cmd.Usage)
		assert.True(t, cmd.Hidden)
		assert.True(t, cmd.Parsed())
	}
}

func TestCommand_Parse_WithNoExecutable(t *testing.T) {
	cmd := &cli.Command{}
	assert.False(t, cmd.Parsed())

	err := cmd.Parse()
	if assert.Error(t, err, "An error was expected") {
		assert.EqualError(t, err, "open : no such file or directory")
		assert.True(t, cmd.Parsed())
	}
}

func TestCommand_Parse_WithNotExistingExecutable(t *testing.T) {
	cmd := &cli.Command{Path: "testdata/foo-not-found"}
	assert.False(t, cmd.Parsed())

	err := cmd.Parse()
	if assert.Error(t, err, "An error was expected") {
		assert.EqualError(t, err, "open testdata/foo-not-found: no such file or directory")
		assert.True(t, cmd.Parsed())
	}
}

func TestCommand_Run_WithAction(t *testing.T) {
	var s string

	cmd := &cli.Command{
		Action: func(ctx *cli.Context) error {
			s = "foo"
			return nil
		},
		Name: "foo",
	}

	err := cmd.Run(&cli.Context{})
	if assert.NoError(t, err) {
		assert.Equal(t, "foo", s)
	}
}

func TestCommand_Run_WithNotExistingExecutable(t *testing.T) {
	cmd := &cli.Command{
		Name: "foo",
		Path: "testdata/foo-not-found",
	}

	err := cmd.Run(&cli.Context{})
	assert.EqualError(t, err, "'foo' appears to be a valid command, but we were not\n"+
		"able to execute it. Maybe foo-not-found is broken?")
}

func TestCommand_Visible(t *testing.T) {
	cmd1 := &cli.Command{Hidden: false}
	cmd2 := &cli.Command{Hidden: true}

	assert.True(t, cmd1.Visible())
	assert.False(t, cmd2.Visible())
}

func TestCommands_Add(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}

	cmds := cli.Commands{cmd1}
	cmds.Add(cmd2)
	assert.Equal(t, 2, cmds.Len())
	assert.Equal(t, cmd1.Name, cmds[0].Name)
	assert.Equal(t, cmd2.Name, cmds[1].Name)
}

func TestCommands_Lookup(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}

	cmds := cli.Commands{cmd1, cmd2}
	assert.Equal(t, cmd1.Name, cmds.Lookup("one").Name)
	assert.Equal(t, cmd2.Name, cmds.Lookup("two").Name)
	assert.Nil(t, cmds.Lookup("three"))
}

func TestCommands_SuggestionsFor(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}
	cmd3 := &cli.Command{Name: "three"}
	cmds := cli.Commands{cmd1, cmd2, cmd3}

	suggestions1 := cmds.SuggestionsFor("neo")
	assert.Equal(t, 1, suggestions1.Len())
	assert.Equal(t, cmd1.Name, suggestions1[0].Name)

	suggestions2 := cmds.SuggestionsFor("t")
	assert.Equal(t, 2, suggestions2.Len())
	assert.Equal(t, cmd3.Name, suggestions2[0].Name)
	assert.Equal(t, cmd2.Name, suggestions2[1].Name)
}

func TestCommands_Visible(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two", Hidden: true}
	cmd3 := &cli.Command{Name: "three"}
	cmds := cli.Commands{cmd1, cmd2, cmd3}

	visible := cmds.Visible()
	assert.Equal(t, 2, visible.Len())
	assert.Equal(t, cmd1.Name, visible[0].Name)
	assert.Equal(t, cmd3.Name, visible[1].Name)
}
