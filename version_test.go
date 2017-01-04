package mondas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testVersionCommand = &VersionCommand{
	name:    "version",
	summary: "version summary",
}

func TestVersionCommand_LoadHelp(t *testing.T) {
	assert.Nil(t, testVersionCommand.LoadHelp())
}

func TestVersionCommand_Name(t *testing.T) {
	assert.Equal(t, "version", testVersionCommand.Name())
}

func TestVersionCommand_ShowHelp(t *testing.T) {
	assert.Nil(t, testVersionCommand.ShowHelp())
}

func TestVersionCommand_Summary(t *testing.T) {
	assert.Equal(t, "version summary", testVersionCommand.Summary())
}

func TestVersionCommand_Usage(t *testing.T) {
	assert.Nil(t, testVersionCommand.Usage())
}
