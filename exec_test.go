package mondas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var execCmd = NewExecCommand("name", "/path/to/file")

func TestNewExecCommand(t *testing.T) {
	assert.Equal(t, "name", execCmd.name)
	assert.Equal(t, "/path/to/file", execCmd.path)
}

func TestExecCommand_Name(t *testing.T) {
	assert.Equal(t, "name", execCmd.Name())
}

func TestExecCommand_Summary(t *testing.T) {
	execCmd.summary = "summary"
	assert.Equal(t, "summary", execCmd.Summary())
}

func TestExecCommand_Usage(t *testing.T) {
	execCmd.usage = []string{"usage"}
	assert.Equal(t, []string{"usage"}, execCmd.Usage())
}
