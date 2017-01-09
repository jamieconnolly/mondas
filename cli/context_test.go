package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	app := &App{}
	args := Args{}

	ctx := NewContext(app, args)
	assert.Equal(t, app, ctx.App)
	assert.Equal(t, args, ctx.Args)
}
