package mondas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var helpCmd = &HelpCommand{
	name:    "help",
	summary: "help summary",
}

func TestHelpCommand_LoadHelp(t *testing.T) {
	assert.Nil(t, helpCmd.LoadHelp())
}

func TestHelpCommand_Name(t *testing.T) {
	assert.Equal(t, "help", helpCmd.Name())
}

func TestHelpCommand_Summary(t *testing.T) {
	assert.Equal(t, "help summary", helpCmd.Summary())
}

func TestHelpCommand_Usage(t *testing.T) {
	assert.Nil(t, helpCmd.Usage())
}
