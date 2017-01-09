package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVersionCommand(t *testing.T) {
	cmd := NewVersionCommand()
	assert.Equal(t, "version", cmd.Name())
	assert.Equal(t, "Display version information", cmd.Summary())
	assert.Empty(t, cmd.Usage())
}

func TestVersionCommand_LoadMetadata(t *testing.T) {
	cmd := NewVersionCommand()
	assert.Nil(t, cmd.LoadMetadata())
}

func TestVersionCommand_ShowHelp(t *testing.T) {
	cmd := NewVersionCommand()
	assert.Nil(t, cmd.ShowHelp())
}
