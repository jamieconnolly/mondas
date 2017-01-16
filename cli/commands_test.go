package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommands_Add(t *testing.T) {
	cmd1 := &Command{Name: "one"}
	cmd2 := &Command{Name: "two"}

	cmds := Commands{cmd1}
	cmds.Add(cmd2)
	assert.Equal(t, 2, cmds.Len())
	assert.Equal(t, cmd1.Name, cmds[0].Name)
	assert.Equal(t, cmd2.Name, cmds[1].Name)
}

func TestCommands_Lookup(t *testing.T) {
	cmd1 := &Command{Name: "one"}
	cmd2 := &Command{Name: "two"}

	cmds := Commands{cmd1, cmd2}
	assert.Equal(t, cmd1.Name, cmds.Lookup("one").Name)
	assert.Equal(t, cmd2.Name, cmds.Lookup("two").Name)
	assert.Nil(t, cmds.Lookup("three"))
}

func TestCommands_SuggestionsFor(t *testing.T) {
	cmd1 := &Command{Name: "one"}
	cmd2 := &Command{Name: "two"}
	cmd3 := &Command{Name: "three"}
	cmds := Commands{cmd1, cmd2, cmd3}

	suggestions1 := cmds.SuggestionsFor("neo")
	assert.Equal(t, 1, suggestions1.Len())
	assert.Equal(t, cmd1.Name, suggestions1[0].Name)

	suggestions2 := cmds.SuggestionsFor("t")
	assert.Equal(t, 2, suggestions2.Len())
	assert.Equal(t, cmd2.Name, suggestions2[0].Name)
	assert.Equal(t, cmd3.Name, suggestions2[1].Name)
}

func TestCommands_Sort(t *testing.T) {
	cmd1 := &Command{Name: "one"}
	cmd2 := &Command{Name: "two"}
	cmd3 := &Command{Name: "three"}

	cmds := (Commands{cmd1, cmd2, cmd3}).Sort()
	assert.Equal(t, cmd1.Name, cmds[0].Name)
	assert.Equal(t, cmd3.Name, cmds[1].Name)
	assert.Equal(t, cmd2.Name, cmds[2].Name)
}
