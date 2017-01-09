package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHelpCommand(t *testing.T) {
	cmd := NewHelpCommand()
	assert.Equal(t, "help", cmd.Name())
	assert.Equal(t, "Display help information", cmd.Summary())
	assert.Empty(t, cmd.Usage())
}

func TestHelpCommand_LoadMetadata(t *testing.T) {
	cmd := NewHelpCommand()
	assert.Nil(t, cmd.LoadMetadata())
}
