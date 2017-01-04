package mondas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExecCommand(t *testing.T) {
	cmd := NewExecCommand("name", "/path/to/file")
	assert.Equal(t, "name", cmd.name)
	assert.Equal(t, "/path/to/file", cmd.path)
}

func TestExecCommand_Name(t *testing.T) {
	cmd := NewExecCommand("name", "/path/to/file")
	assert.Equal(t, "name", cmd.Name())
}

func TestExecCommand_Summary(t *testing.T) {
	cmd := NewExecCommand("name", "/path/to/file")
	cmd.summary = "summary"
	assert.Equal(t, "summary", cmd.Summary())
}

func TestExecCommand_Usage(t *testing.T) {
	cmd := NewExecCommand("name", "/path/to/file")
	cmd.usage = []string{"usage"}
	assert.Equal(t, []string{"usage"}, cmd.Usage())
}
