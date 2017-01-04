package mondas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var versionCmd = &VersionCommand{
	name:    "version",
	summary: "version summary",
}

func TestVersionCommand_LoadHelp(t *testing.T) {
	assert.Nil(t, versionCmd.LoadHelp())
}

func TestVersionCommand_Name(t *testing.T) {
	assert.Equal(t, "version", versionCmd.Name())
}

func TestVersionCommand_ShowHelp(t *testing.T) {
	assert.Nil(t, versionCmd.ShowHelp())
}

func TestVersionCommand_Summary(t *testing.T) {
	assert.Equal(t, "version summary", versionCmd.Summary())
}

func TestVersionCommand_Usage(t *testing.T) {
	assert.Nil(t, versionCmd.Usage())
}
