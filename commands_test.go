package mondas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCommand struct {
	name string
}

func (c *testCommand) LoadHelp() error        { return nil }
func (c *testCommand) Name() string           { return c.name }
func (c *testCommand) Run(ctx *Context) error { return nil }
func (c *testCommand) ShowHelp() error        { return nil }
func (c *testCommand) Summary() string        { return "" }
func (c *testCommand) Usage() []string        { return nil }

var (
	barCommand = &testCommand{name: "bar"}
	fooCommand = &testCommand{name: "foo"}
)

func TestCommands_Add(t *testing.T) {
	cmds := Commands{fooCommand}

	cmds.Add(barCommand)
	cmds.Add(fooCommand)
	assert.Equal(t, 2, cmds.Len())
	assert.Equal(t, fooCommand, cmds[0])
	assert.Equal(t, barCommand, cmds[1])
}

func TestCommands_Lookup(t *testing.T) {
	cmds := Commands{fooCommand, barCommand}
	assert.Equal(t, fooCommand, cmds.Lookup("foo"))
	assert.Equal(t, barCommand, cmds.Lookup("bar"))
	assert.Nil(t, cmds.Lookup("baz"))
}

func TestCommands_Sort(t *testing.T) {
	cmds := Commands{fooCommand, barCommand}

	cmds.Sort()
	assert.Equal(t, barCommand, cmds[0])
	assert.Equal(t, fooCommand, cmds[1])
}
