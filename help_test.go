package mondas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testHelpCommand = &HelpCommand{
	name:    "help",
	summary: "help summary",
}

func TestHelpCommand_LoadHelp(t *testing.T) {
	assert.Nil(t, testHelpCommand.LoadHelp())
}

func TestHelpCommand_Name(t *testing.T) {
	assert.Equal(t, "help", testHelpCommand.Name())
}

func TestHelpCommand_Summary(t *testing.T) {
	assert.Equal(t, "help summary", testHelpCommand.Summary())
}

func TestHelpCommand_Usage(t *testing.T) {
	assert.Nil(t, testHelpCommand.Usage())
}
