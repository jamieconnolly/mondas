package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExecCommand(t *testing.T) {
	cmd := NewExecCommand("name", "/path/to/file")
	assert.Equal(t, "name", cmd.Name())
	assert.Equal(t, "/path/to/file", cmd.Path())
	assert.Empty(t, cmd.Summary())
	assert.Empty(t, cmd.Usage())
}
