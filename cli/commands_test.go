package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeCommand struct {
	name string
}

func (c *FakeCommand) LoadMetadata() error    { return nil }
func (c *FakeCommand) Name() string           { return c.name }
func (c *FakeCommand) Run(ctx *Context) error { return nil }
func (c *FakeCommand) ShowHelp() error        { return nil }
func (c *FakeCommand) Summary() string        { return "" }
func (c *FakeCommand) Usage() []string        { return nil }

func TestCommands_Add(t *testing.T) {
	cmd1 := &FakeCommand{name: "one"}
	cmd2 := &FakeCommand{name: "two"}

	cmds := Commands{cmd1}
	cmds.Add(cmd2)
	assert.Equal(t, 2, cmds.Len())
	assert.Equal(t, cmd1, cmds[0])
	assert.Equal(t, cmd2, cmds[1])
}

func TestCommands_Lookup(t *testing.T) {
	cmd1 := &FakeCommand{name: "one"}
	cmd2 := &FakeCommand{name: "two"}

	cmds := Commands{cmd1, cmd2}
	assert.Equal(t, cmd1, cmds.Lookup("one"))
	assert.Equal(t, cmd2, cmds.Lookup("two"))
	assert.Nil(t, cmds.Lookup("three"))
}

func TestCommands_SuggestionsFor(t *testing.T) {
	cmd1 := &FakeCommand{name: "one"}
	cmd2 := &FakeCommand{name: "two"}
	cmd3 := &FakeCommand{name: "three"}
	cmds := Commands{cmd1, cmd2, cmd3}

	suggestions1 := cmds.SuggestionsFor("neo")
	assert.Equal(t, 1, suggestions1.Len())
	assert.Equal(t, cmd1, suggestions1[0])

	suggestions2 := cmds.SuggestionsFor("t")
	assert.Equal(t, 2, suggestions2.Len())
	assert.Equal(t, cmd3, suggestions2[0])
	assert.Equal(t, cmd2, suggestions2[1])
}

func TestCommands_Sort(t *testing.T) {
	cmd1 := &FakeCommand{name: "one"}
	cmd2 := &FakeCommand{name: "two"}
	cmd3 := &FakeCommand{name: "three"}

	cmds := Commands{cmd1, cmd2, cmd3}
	cmds.Sort()
	assert.Equal(t, cmd1, cmds[0])
	assert.Equal(t, cmd3, cmds[1])
	assert.Equal(t, cmd2, cmds[2])
}
