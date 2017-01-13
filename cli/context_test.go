package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	app := &App{}
	args := []string{"foo"}
	env := []string{"foo=bar", "baz=foo"}

	ctx := NewContext(app, args, env)
	assert.Equal(t, app, ctx.App)
	assert.Equal(t, Args(args), ctx.Args)
	assert.Equal(t, NewEnvFromEnviron(env), ctx.Env)
}
